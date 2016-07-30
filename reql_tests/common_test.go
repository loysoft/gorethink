package reql_tests

import "testing"

func TestCompareString(t *testing.T) {

	// simple
	assertCompare(t, "a", "a")
	assertCompare(t, "á", "á")
	assertCompare(t, "something longer\nwith two lines", "something longer\nwith two lines")

	assertCompareFalse(t, "a", "b")
	assertCompareFalse(t, "a", 1)
	assertCompareFalse(t, "a", []interface{}{})
	assertCompareFalse(t, "a", nil)
	assertCompareFalse(t, "a", []interface{}{"a"})
	assertCompareFalse(t, "a", map[string]interface{}{"a": 1})
}
func TestCompareArray(t *testing.T) {

	// simple pass
	assertCompare(t, []interface{}{1, 2, 3}, []interface{}{1, 2, 3})

	// out of order
	assertCompareFalse(t, []interface{}{1, 2, 3}, []interface{}{1, 3, 2})

	// totally mistmatched lists
	assertCompareFalse(t, []interface{}{1, 2, 3}, []interface{}{3, 4, 5})

	// missing items
	assertCompareFalse(t, []interface{}{1, 2, 3}, []interface{}{1, 2})
	assertCompareFalse(t, []interface{}{1, 2, 3}, []interface{}{1, 3})

	// extra items
	assertCompareFalse(t, []interface{}{1, 2, 3}, []interface{}{1, 2, 3, 4})

	// empty array
	assertCompare(t, []interface{}{}, []interface{}{})
	assertCompareFalse(t, []interface{}{1}, []interface{}{})
	assertCompareFalse(t, []interface{}{}, []interface{}{1})
	assertCompareFalse(t, []interface{}{}, nil)

	// strings
	assertCompare(t, []interface{}{"a", "b"}, []interface{}{"a", "b"})
	assertCompareFalse(t, []interface{}{"a", "c"}, []interface{}{"a", "b"})

	// multiple of a single value
	assertCompare(t, []interface{}{1, 2, 2, 3, 3, 3}, []interface{}{1, 2, 2, 3, 3, 3})
	assertCompareFalse(t, []interface{}{1, 2, 2, 3, 3, 3}, []interface{}{1, 2, 3})
	assertCompareFalse(t, []interface{}{1, 2, 3}, []interface{}{1, 2, 2, 3, 3, 3})
}
func TestCompareArray_partial(t *testing.T) {
	// note that these are all in-order

	// simple
	assertCompare(t, partial([]interface{}{1}), []interface{}{1, 2, 3})
	assertCompare(t, partial([]interface{}{2}), []interface{}{1, 2, 3})
	assertCompare(t, partial([]interface{}{3}), []interface{}{1, 2, 3})

	assertCompare(t, partial([]interface{}{1, 2}), []interface{}{1, 2, 3})
	assertCompare(t, partial([]interface{}{1, 3}), []interface{}{1, 2, 3})

	assertCompare(t, partial([]interface{}{1, 2, 3}), []interface{}{1, 2, 3})

	assertCompareFalse(t, partial([]interface{}{4}), []interface{}{1, 2, 3})

	// ordered
	assertCompareFalse(t, partial([]interface{}{3, 2, 1}).SetOrdered(true), []interface{}{1, 2, 3})
	assertCompareFalse(t, partial([]interface{}{1, 3, 2}).SetOrdered(true), []interface{}{1, 2, 3})

	// empty array
	assertCompare(t, partial([]interface{}{}), []interface{}{1, 2, 3})

	// multiple of a single items
	assertCompare(t, partial([]interface{}{1, 2, 2}), []interface{}{1, 2, 2, 3, 3, 3})
	assertCompareFalse(t, partial([]interface{}{1, 2, 2, 2}), []interface{}{1, 2, 2, 3, 3, 3})
}
func TestCompareArray_unordered(t *testing.T) {

	// simple
	assertCompare(t, bag([]interface{}{1, 2}), []interface{}{1, 2})
	assertCompare(t, bag([]interface{}{2, 1}), []interface{}{1, 2})

	assertCompareFalse(t, bag([]interface{}{1, 2}), []interface{}{1, 2, 3})
	assertCompareFalse(t, bag([]interface{}{1, 3}), []interface{}{1, 2, 3})
	assertCompareFalse(t, bag([]interface{}{3, 1}), []interface{}{1, 2, 3})

	// empty array
	assertCompare(t, bag([]interface{}{}), []interface{}{})
}
func TestCompareMap(t *testing.T) {

	// simple
	assertCompare(t, map[string]interface{}{"a": 1, "b": 2, "c": 3}, map[string]interface{}{"a": 1, "b": 2, "c": 3})
	assertCompare(t, map[string]interface{}{"a": 1, "b": 2, "c": 3}, map[string]interface{}{"c": 3, "a": 1, "b": 2})

	assertCompareFalse(t, map[string]interface{}{"a": 1, "b": 2, "c": 3}, map[string]interface{}{"a": 1})
	assertCompareFalse(t, map[string]interface{}{"a": 1}, map[string]interface{}{"a": 1, "b": 2, "c": 3})

	// empty
	assertCompare(t, map[string]interface{}{}, map[string]interface{}{})
	assertCompareFalse(t, map[string]interface{}{}, map[string]interface{}{"a": 1})
	assertCompareFalse(t, map[string]interface{}{"a": 1}, map[string]interface{}{})

	assertCompare(t, map[interface{}]interface{}{1: 1225, 2: 1250, 3: 1275, 0: 1200}, map[string]interface{}{"2": 1250, "3": 1275, "0": 1200, "1": 1225})
	assertCompare(t, map[interface{}]interface{}{nil: 33, 0: 22, 20: 22, 30: 23}, map[string]interface{}{"30": 23, "": 33, "0": 22, "20": 22})
}
func TestCompareMap_partial(t *testing.T) {

	// simple
	assertCompare(t, partial(map[string]interface{}{"a": 1}), map[string]interface{}{"a": 1})
	assertCompare(t, partial(map[string]interface{}{"a": 1}), map[string]interface{}{"a": 1, "b": 2})

	assertCompareFalse(t, partial(map[string]interface{}{"a": 2}), map[string]interface{}{"a": 1, "b": 2})
	assertCompareFalse(t, partial(map[string]interface{}{"c": 1}), map[string]interface{}{"a": 1, "b": 2})
	assertCompareFalse(t, partial(map[string]interface{}{"a": 1, "b": 2}), map[string]interface{}{"b": 2})

	// empty
	assertCompare(t, partial(map[string]interface{}{}), map[string]interface{}{})
	assertCompare(t, partial(map[string]interface{}{}), map[string]interface{}{"a": 1})
	assertCompareFalse(t, partial(map[string]interface{}{"a": 1}), map[string]interface{}{})
}
func TestCompareMap_inSlice(t *testing.T) {

	// simple
	assertCompare(t, []interface{}{map[string]interface{}{"a": 1}}, []interface{}{map[string]interface{}{"a": 1}})
	assertCompare(t, []interface{}{map[string]interface{}{"a": 1, "b": 2}}, []interface{}{map[string]interface{}{"a": 1, "b": 2}})
	assertCompare(t, []interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}}, []interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}})

	assertCompareFalse(t, []interface{}{map[string]interface{}{"a": 1}}, []interface{}{map[string]interface{}{"a": 1, "b": 2}})
	assertCompareFalse(t, []interface{}{map[string]interface{}{"a": 2, "b": 2}}, []interface{}{map[string]interface{}{"a": 1, "b": 2}})
	assertCompareFalse(t, []interface{}{map[string]interface{}{"a": 2, "c": 3}}, []interface{}{map[string]interface{}{"a": 1, "b": 2}})
	assertCompareFalse(t, []interface{}{map[string]interface{}{"a": 2, "c": 3}}, []interface{}{map[string]interface{}{"a": 1}})
	assertCompareFalse(t, []interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}}, []interface{}{map[string]interface{}{"a": 1, "b": 2}})

	// order
	assertCompareFalse(t, []interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}}, []interface{}{map[string]interface{}{"b": 2}, map[string]interface{}{"a": 1}})

	// partial
	assertCompare(t, partial([]interface{}{map[string]interface{}{}}), []interface{}{map[string]interface{}{"a": 1, "b": 2}})
	assertCompare(t, partial([]interface{}{map[string]interface{}{}}), []interface{}{map[string]interface{}{"a": 1, "b": 2}})
	assertCompare(t, partial([]interface{}{map[string]interface{}{"a": 1}}), []interface{}{map[string]interface{}{"a": 1, "b": 2}})
	assertCompare(t, partial([]interface{}{map[string]interface{}{"a": 1, "b": 2}}), []interface{}{map[string]interface{}{"a": 1, "b": 2}})
	assertCompare(t, partial([]interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}}), []interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}, map[string]interface{}{"c": 3}})

	assertCompareFalse(t, partial([]interface{}{map[string]interface{}{"a": 2}}), []interface{}{map[string]interface{}{"a": 1, "b": 2}})
	assertCompareFalse(t, partial([]interface{}{map[string]interface{}{"a": 1, "b": 2}}), []interface{}{map[string]interface{}{"a": 1}})

	// partial order
	assertCompareFalse(t, partial([]interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}}).SetOrdered(true), []interface{}{map[string]interface{}{"b": 2}, map[string]interface{}{"a": 1}})

	// partial unordered
	assertCompare(t, partial([]interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}}), []interface{}{map[string]interface{}{"b": 2}, map[string]interface{}{"a": 1}})
	assertCompare(t, partial([]interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}}).SetOrdered(false), []interface{}{map[string]interface{}{"b": 2}, map[string]interface{}{"a": 1}})

	assertCompare(t, []interface{}{map[string]interface{}{"a": 1, "b": 1}, partial(map[string]interface{}{"a": 2})}, []interface{}{map[string]interface{}{"a": 1, "b": 1}, map[string]interface{}{"a": 2, "b": 2}})
}

func TestCompareUUID(t *testing.T) {

	// simple
	assertCompare(t, uuid(), "4e9e5bc2-9b11-4143-9aa1-75c10e7a193a")
	assertCompareFalse(t, uuid(), "4")
	assertCompareFalse(t, uuid(), "*")
	assertCompareFalse(t, uuid(), nil)
}

func TestCompareNumbers(t *testing.T) {

	// simple
	assertCompare(t, 1, 1)
	assertCompare(t, 1, 1.0)
	assertCompare(t, 1.0, 1)
	assertCompare(t, 1.0, 1.0)

	assertCompareFalse(t, 1, 2)
	assertCompareFalse(t, 1, 2.0)
	assertCompareFalse(t, 1.0, 2)
	assertCompareFalse(t, 1.0, 2.0)

	// precision
	assertComparePrecision(t, 1, 1.4, 0.5)
	assertComparePrecision(t, 1.0, 1.4, 0.5)

	assertComparePrecisionFalse(t, 1, 2, 0.5)
	assertComparePrecisionFalse(t, 1, 1.6, 0.5)
	assertComparePrecisionFalse(t, 1.0, 2, 0.5)
	assertComparePrecisionFalse(t, 1.0, 1.6, 0.5)
}
