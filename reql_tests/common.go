package reql_tests

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	r "gopkg.in/dancannon/gorethink.v2"
)

func maybeRun(query interface{}, session *r.Session, opts r.RunOpts) interface{} {
	switch query := query.(type) {
	case r.Term:
		cursor, err := query.Run(session, opts)
		if err != nil {
			return err
		}

		switch cursor.Type() {
		case "Cursor":
			if cursor.IsSingleValue() {
				var result interface{}
				if err := cursor.One(&result); err != nil {
					return err
				}
				return result
			} else {
				var results []interface{}
				if err := cursor.All(&results); err != nil {
					return err
				}
				return results
			}
		default:
			// If this is a changefeed then return the cursor without attempting
			// to read any documents
			return cursor
		}
	default:
		return query
	}
}

func runAndAssert(suite suite.Suite, expected, v interface{}, session *r.Session, opts r.RunOpts) {
	var cursor *r.Cursor
	var err error

	switch v := v.(type) {
	case r.Term:
		cursor, err = v.Run(session, opts)
	case *r.Cursor:
		cursor = v
	case error:
		err = v
	}

	assertExpected(suite, expected, cursor, err)
}

func fetchAndAssert(suite suite.Suite, expected, result interface{}, count int) {
	switch v := expected.(type) {
	case Expected:
		v.Fetch = true
		v.FetchCount = count

		expected = v
	default:
		expected = Expected{
			Val:        v,
			Fetch:      true,
			FetchCount: count,
		}
	}

	var cursor *r.Cursor
	var err error

	switch result := result.(type) {
	case *r.Cursor:
		cursor = result
	case error:
		err = result
	}

	assertExpected(suite, expected, cursor, err)
}

func maybeLen(v interface{}) interface{} {
	switch v := v.(type) {
	case *r.Cursor:
		results := []interface{}{}
		v.All(&results)
		return len(results)
	case []interface{}:
		return len(v)
	default:
		return v
	}
}

func assertExpected(suite suite.Suite, expected interface{}, obtainedCursor *r.Cursor, obtainedErr error) {
	if expected == AnythingIsFine {
		suite.NoError(obtainedErr, "Query returned unexpected error")
		return
	}

	switch expected := expected.(type) {
	case Err:
		expected.assert(suite, obtainedCursor, obtainedErr)
	case Expected:
		expected.assert(suite, obtainedCursor, obtainedErr)
	default:
		Expected{Val: expected}.assert(suite, obtainedCursor, obtainedErr)
	}
}

type Expected struct {
	Fetch      bool
	Partial    bool
	Ordered    bool
	FetchCount int
	Val        interface{}
}

func (expected Expected) SetOrdered(ordered bool) Expected {
	expected.Ordered = ordered

	return expected
}

func (expected Expected) SetPartial(partial bool) Expected {
	expected.Partial = partial

	return expected
}

func (expected Expected) assert(suite suite.Suite, obtainedCursor *r.Cursor, obtainedErr error) {
	if suite.NoError(obtainedErr, "Query returned unexpected error") {
		return
	}

	expectedVal := reflect.ValueOf(expected.Val)

	// If expected value is nil then ensure cursor is nil (assume that an
	// invalid reflect value is because expected value is nil)
	if !expectedVal.IsValid() || (expectedVal.Kind() == reflect.Ptr && expectedVal.IsNil()) {
		suite.True(obtainedCursor.IsNil(), "Expected nil cursor")
		return
	}

	expectedType := expectedVal.Type()
	expectedKind := expectedType.Kind()

	if expectedKind == reflect.Array || expectedKind == reflect.Slice || expected.Fetch {
		if expectedType.Elem().Kind() == reflect.Uint8 {
			// Decode byte slices slightly differently
			var obtained = []byte{}
			err := obtainedCursor.One(&obtained)
			suite.NoError(err, "Error returned when reading query response")
			assertCompare(suite.T(), expected, obtained)
		} else {
			var obtained = []interface{}{}
			if expected.Fetch {
				var v interface{}
				for obtainedCursor.Next(&v) {
					obtained = append(obtained, v)

					if expected.FetchCount != 0 && len(obtained) >= expected.FetchCount {
						break
					}
				}
				suite.NoError(obtainedCursor.Err(), "Error returned when reading query response")
			} else {
				err := obtainedCursor.All(&obtained)
				suite.NoError(err, "Error returned when reading query response")
			}

			assertCompare(suite.T(), expected, obtained)
		}
	} else if expectedKind == reflect.Map {
		var obtained map[string]interface{}
		err := obtainedCursor.One(&obtained)
		suite.NoError(err, "Error returned when reading query response")
		assertCompare(suite.T(), expected, obtained)
	} else {
		var obtained interface{}
		err := obtainedCursor.One(&obtained)
		suite.NoError(err, "Error returned when reading query response")
		assertCompare(suite.T(), expected, obtained)
	}
}

func assertCompare(t *testing.T, expected, actual interface{}) {
	expectedVal := expected
	if e, ok := expected.(Expected); ok {
		expectedVal = e.Val
	}

	ok, msg := compare(expected, actual)
	if !ok {
		assert.Fail(t, fmt.Sprintf("Not equal: %#v (expected)\n           != %#v (actual)", expectedVal, actual), msg)
	}
}

func assertCompareFalse(t *testing.T, expected, actual interface{}) {
	expectedVal := expected
	if e, ok := expected.(Expected); ok {
		expectedVal = e.Val
	}

	ok, msg := compare(expected, actual)
	if ok {
		assert.Fail(t, fmt.Sprintf("Should not be equal: %#v (expected)\n           == %#v (actual)", expectedVal, actual), msg)
	}
}

func assertComparePrecision(t *testing.T, expected, actual interface{}, precision float64) {
	expectedVal := expected
	if e, ok := expected.(Expected); ok {
		expectedVal = e.Val
	}

	ok, msg := comparePrecision(expected, actual, precision)
	if !ok {
		assert.Fail(t, fmt.Sprintf("Not equal: %#v (expected)\n           != %#v (actual)", expectedVal, actual), msg)
	}
}

func assertComparePrecisionFalse(t *testing.T, expected, actual interface{}, precision float64) {
	expectedVal := expected
	if e, ok := expected.(Expected); ok {
		expectedVal = e.Val
	}

	ok, msg := comparePrecision(expected, actual, precision)
	if ok {
		assert.Fail(t, fmt.Sprintf("Should not be equal: %#v (expected)\n           == %#v (actual)", expectedVal, actual), msg)
	}
}

func compare(expected, actual interface{}) (bool, string) {
	return comparePrecision(expected, actual, 0.00000000001)
}

func comparePrecision(expected, actual interface{}, precision float64) (bool, string) {
	return compareOpts(expected, actual, true, false, precision)
}

func compareOpts(expected, actual interface{}, ordered, partial bool, precision float64) (bool, string) {
	if e, ok := expected.(Expected); ok {
		partial = e.Partial
		ordered = e.Ordered
		expected = e.Val
	}

	// Anything
	if expected == AnythingIsFine {
		return true, ""
	}

	expectedVal := reflect.ValueOf(expected)
	actualVal := reflect.ValueOf(actual)

	// Nil
	if expected == nil {
		switch actualVal.Kind() {
		case reflect.Bool:
			expected = false
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			expected = 0.0
		case reflect.String:
			expected = ""
		}

		if expected == actual {
			return true, ""
		}
	}

	// Regex
	if expr, ok := expected.(Regex); ok {
		re, err := regexp.Compile(string(expr))
		if err != nil {
			return false, fmt.Sprintf("Failed to compile regexp: %s", err)
		}

		if actualVal.Kind() != reflect.String {
			return false, fmt.Sprintf("Expected string, got %t (%T)", actual, actual)
		}

		if !re.MatchString(actualVal.String()) {
			return false, fmt.Sprintf("Value %v did not match regexp '%s'", actual, expr)
		}

		return true, ""
	}

	switch expectedVal.Kind() {

	// Bool
	case reflect.Bool:
		if expected == actual {
			return true, ""
		}
	// Number
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		switch actualVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64, reflect.String:
			diff := math.Abs(reflectNumber(expectedVal) - reflectNumber(actualVal))
			if diff <= precision {
				return true, ""
			}

			if precision != 0 {
				return false, fmt.Sprintf("Value %v was not within %f of %v", expected, precision, actual)
			}

			return false, fmt.Sprintf("Expected %v but got %v", expected, actual)
		}

	// String
	case reflect.String:
		actualStr := fmt.Sprintf("%v", actual)
		if expected == actualStr {
			return true, ""
		}
	// Struct
	case reflect.Struct:
		// Convert expected struct to map and compare with actual value
		return compareOpts(reflectMap(expectedVal), actual, ordered, partial, precision)
	// Map
	case reflect.Map:
		switch actualVal.Kind() {
		case reflect.Struct:
			// Convert actual struct to map and compare with expected map
			return compareOpts(expected, reflectMap(actualVal), ordered, partial, precision)
		case reflect.Map:
			expectedKeys := expectedVal.MapKeys()
			actualKeys := actualVal.MapKeys()

			for _, expectedKey := range expectedKeys {
				keyFound := false
				for _, actualKey := range actualKeys {
					if ok, _ := compare(expectedKey.Interface(), actualKey.Interface()); ok {
						keyFound = true
						break
					}
				}
				if !keyFound {
					return false, fmt.Sprintf("Expected field %v but not found", expectedKey)
				}
			}

			if !partial {
				expectedKeyVals := reflectMapKeys(expectedKeys)
				actualKeyVals := reflectMapKeys(actualKeys)
				if ok, _ := compareOpts(expectedKeyVals, actualKeyVals, false, false, 0.0); !ok {
					return false, fmt.Sprintf(
						"Unmatched keys from either side: expected fields %v, got %v",
						expectedKeyVals, actualKeyVals,
					)
				}
			}

			expectedMap := reflectMap(expectedVal)
			actualMap := reflectMap(actualVal)

			for k, v := range expectedMap {
				if ok, reason := compareOpts(v, actualMap[k], ordered, partial, precision); !ok {
					return false, reason
				}
			}

			return true, ""
		default:
			return false, fmt.Sprintf("Expected map, got %v (%T)", actual, actual)
		}
	// Slice/Array
	case reflect.Slice, reflect.Array:
		switch actualVal.Kind() {
		case reflect.Slice, reflect.Array:
			if ordered {
				expectedArr := reflectSlice(expectedVal)
				actualArr := reflectSlice(actualVal)

				j := 0
				for i := 0; i < len(expectedArr); i++ {
					expectedArrVal := expectedArr[i]
					for {
						if j >= len(actualArr) {
							return false, fmt.Sprintf("Ran out of results before finding %v", expectedArrVal)
						}

						actualArrVal := actualArr[j]
						j++

						if ok, _ := compareOpts(expectedArrVal, actualArrVal, ordered, partial, precision); ok {
							break
						} else if !partial {
							return false, fmt.Sprintf("Unexpected item %v while looking for %v", actualArrVal, expectedArrVal)
						}
					}
				}
				if !partial && j < len(actualArr) {
					return false, fmt.Sprintf("Unexpected extra results: %v", actualArr[j:])
				}
			} else {
				expectedArr := reflectSlice(expectedVal)
				actualArr := reflectSlice(actualVal)

				for _, expectedArrVal := range expectedArr {
					found := false
					for j, actualArrVal := range actualArr {
						if ok, _ := compareOpts(expectedArrVal, actualArrVal, ordered, partial, precision); ok {
							found = true
							actualArr = append(actualArr[:j], actualArr[j+1:]...)
							break
						}
					}
					if !found {
						return false, fmt.Sprintf("Missing expected item %v", expectedArrVal)
					}
				}

				if !partial && len(actualArr) > 0 {
					return false, fmt.Sprintf("Extra items returned: %v", expectedArr)
				}
			}

			return true, ""
		}
	// Other
	default:
		if expected == actual {
			return true, ""
		}
	}

	return false, fmt.Sprintf("Expected %v (%T) but got %v (%T)", expected, expected, actual, actual)
}

func reflectNumber(v reflect.Value) float64 {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint())
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.String:
		f, _ := strconv.ParseFloat(v.String(), 64)
		return f
	default:
		return float64(0)
	}
}

func reflectMap(v reflect.Value) map[interface{}]interface{} {
	switch v.Kind() {
	case reflect.Struct:
		m := map[interface{}]interface{}{}
		for i := 0; i < v.NumField(); i++ {
			sf := v.Type().Field(i)
			if sf.PkgPath != "" && !sf.Anonymous {
				continue // unexported
			}

			k := sf.Name
			v := v.Field(i).Interface()

			m[k] = v
		}
		return m
	case reflect.Map:
		m := map[interface{}]interface{}{}
		for _, mk := range v.MapKeys() {
			k := ""
			if mk.Interface() != nil {
				k = fmt.Sprintf("%v", mk.Interface())
			}
			v := v.MapIndex(mk).Interface()

			m[k] = v
		}
		return m
	default:
		return nil
	}
}

func reflectSlice(v reflect.Value) []interface{} {
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		s := []interface{}{}
		for i := 0; i < v.Len(); i++ {
			s = append(s, v.Index(i).Interface())
		}
		return s
	default:
		return nil
	}
}

func reflectMapKeys(keys []reflect.Value) []interface{} {
	s := []interface{}{}
	for _, key := range keys {
		s = append(s, key.Interface())
	}
	return s
}

func reflectInterfaces(vals []reflect.Value) []interface{} {
	ret := []interface{}{}
	for _, val := range vals {
		ret = append(ret, val.Interface())
	}
	return ret
}

func deleteFromSlice(v []interface{}, i int) []interface{} {
	return append(v[:i], v[i+1:]...)
}

func int_cmp(i int) int {
	return i
}

func float_cmp(i float64) float64 {
	return i
}

func bag(v interface{}) Expected {
	return Expected{
		Ordered: false,
		Partial: false,
		Val:     v,
	}
}

func partial(v interface{}) Expected {
	return Expected{
		Ordered: false,
		Partial: true,
		Val:     v,
	}
}

func arrlen(length int, vs ...interface{}) []interface{} {
	var v interface{} = AnythingIsFine
	if len(vs) == 1 {
		v = vs[0]
	}

	arr := make([]interface{}, length)
	for i := 0; i < length; i++ {
		arr[i] = v
	}
	return arr
}

func str(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func wait(s int) interface{} {
	time.Sleep(time.Duration(s) * time.Second)

	return nil
}

type Err struct {
	Type    string
	Message string
	Regex   string
}

var exceptionRegex = regexp.MustCompile("^(?P<message>[^\n]*?)((?: in:)?\n|\nFailed assertion:)(?s).*$")

func (expected Err) assert(suite suite.Suite, obtainerCursor *r.Cursor, obtainedErr error) {
	// If the error is nil then attempt to read from the cursor and see if an
	// error is returned
	if obtainedErr == nil {
		var res []interface{}
		obtainedErr = obtainerCursor.All(&res)
	}

	if suite.Error(obtainedErr) {
		return
	}

	obtainedType := reflect.TypeOf(obtainedErr).String()
	obtainedMessage := strings.TrimPrefix(obtainedErr.Error(), "gorethink: ")
	obtainedMessage = exceptionRegex.ReplaceAllString(obtainedMessage, "${message}")

	suite.Equal(expected.Type, obtainedType)
	if expected.Regex != "" {
		suite.Regexp(expected.Regex, obtainedMessage)
	}
	if expected.Message != "" {
		suite.Equal(expected.Message, obtainedMessage)
	}
}

func err(errType, message string) Err {
	return Err{
		Type:    "gorethink.RQL" + errType[4:],
		Message: message,
	}
}

func err_regex(errType, expr string) Err {
	return Err{
		Type:  "gorethink.RQL" + errType[4:],
		Regex: expr,
	}
}

type Regex string

func uuid() Regex {
	return Regex("[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}")
}

func regex(expr string) Regex {
	return Regex(expr)
}

var AnythingIsFine = "reql_test.AnythingIsFine"

var Ast = struct {
	RqlTzinfo     func(tz string) *time.Location
	Fromtimestamp func(ts float64, loc *time.Location) time.Time
	Now           func() time.Time
}{
	func(tz string) *time.Location {
		t, _ := time.Parse("-07:00 UTC", tz+" UTC")

		return t.Location()
	},
	func(ts float64, loc *time.Location) time.Time {
		sec, nsec := math.Modf(ts)

		return time.Unix(int64(sec), int64(nsec*1000)*1000000).In(loc)
	},
	time.Now,
}

func UTCTimeZone() *time.Location {
	return time.UTC
}

func PacificTimeZone() *time.Location {
	return Ast.RqlTzinfo("-07:00")
}

var FloatInfo = struct {
	Min, Max float64
}{math.SmallestNonzeroFloat64, math.MaxFloat64}

var sys = struct {
	FloatInfo struct {
		Min, Max float64
	}
}{FloatInfo}
