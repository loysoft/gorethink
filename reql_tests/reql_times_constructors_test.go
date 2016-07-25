// Autogenerated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../../gen_tests/template.go
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/dancannon/gorethink.v2"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTimesConstructorsSuite(t *testing.T) {
    suite.Run(t, new(TimesConstructorsSuite ))
}

type TimesConstructorsSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TimesConstructorsSuite) SetupTest() {
	// Use imports to prevent errors
	time.Now()

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

    r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *TimesConstructorsSuite) TearDownSuite() {
	r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
    r.DBDrop("test").Exec(suite.session)

    suite.session.Close()
}

func (suite *TimesConstructorsSuite) TestCases() {


    {
        // times/constructors.yaml line #5
        /* datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00')) */
        var expected_ time.Time = Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00"))
        /* r.expr(r.epoch_time(896571240)) */

    	suite.T().Log("About to run line #5: r.Expr(r.EpochTime(896571240))")

        cursor, err := r.Expr(r.EpochTime(896571240)).Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #5")
    }

    {
        // times/constructors.yaml line #11
        /* {'stuff':datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00')), 'more':[datetime.fromtimestamp(996571240, r.ast.RqlTzinfo('00:00'))]} */
        var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"stuff": Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00")), "more": []interface{}{Ast.Fromtimestamp(996571240, Ast.RqlTzinfo("00:00"))}, }
        /* r.expr({'stuff':r.epoch_time(896571240), 'more':[r.epoch_time(996571240)]}) */

    	suite.T().Log("About to run line #11: r.Expr(map[interface{}]interface{}{'stuff': r.EpochTime(896571240), 'more': []interface{}{r.EpochTime(996571240)}, })")

        cursor, err := r.Expr(map[interface{}]interface{}{"stuff": r.EpochTime(896571240), "more": []interface{}{r.EpochTime(996571240)}, }).Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #11")
    }

    {
        // times/constructors.yaml line #17
        /* [datetime.fromtimestamp(796571240, r.ast.RqlTzinfo('00:00')), datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00')), {'stuff':datetime.fromtimestamp(996571240, r.ast.RqlTzinfo('00:00'))}] */
        var expected_ []interface{} = []interface{}{Ast.Fromtimestamp(796571240, Ast.RqlTzinfo("00:00")), Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00")), map[interface{}]interface{}{"stuff": Ast.Fromtimestamp(996571240, Ast.RqlTzinfo("00:00")), }}
        /* r.expr([r.epoch_time(796571240), r.epoch_time(896571240), {'stuff':r.epoch_time(996571240)}]) */

    	suite.T().Log("About to run line #17: r.Expr([]interface{}{r.EpochTime(796571240), r.EpochTime(896571240), map[interface{}]interface{}{'stuff': r.EpochTime(996571240), }})")

        cursor, err := r.Expr([]interface{}{r.EpochTime(796571240), r.EpochTime(896571240), map[interface{}]interface{}{"stuff": r.EpochTime(996571240), }}).Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #17")
    }

    {
        // times/constructors.yaml line #23
        /* {'nested':{'time':datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00'))}} */
        var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"nested": map[interface{}]interface{}{"time": Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00")), }, }
        /* r.expr({'nested':{'time':r.epoch_time(896571240)}}) */

    	suite.T().Log("About to run line #23: r.Expr(map[interface{}]interface{}{'nested': map[interface{}]interface{}{'time': r.EpochTime(896571240), }, })")

        cursor, err := r.Expr(map[interface{}]interface{}{"nested": map[interface{}]interface{}{"time": r.EpochTime(896571240), }, }).Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #23")
    }

    {
        // times/constructors.yaml line #29
        /* [1, "two", ["a", datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00')), 3]] */
        var expected_ []interface{} = []interface{}{1, "two", []interface{}{"a", Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00")), 3}}
        /* r.expr([1, "two", ["a", r.epoch_time(896571240), 3]]) */

    	suite.T().Log("About to run line #29: r.Expr([]interface{}{1, 'two', []interface{}{'a', r.EpochTime(896571240), 3}})")

        cursor, err := r.Expr([]interface{}{1, "two", []interface{}{"a", r.EpochTime(896571240), 3}}).Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #29")
    }

    {
        // times/constructors.yaml line #35
        /* 1 */
        var expected_ int = 1
        /* r.epoch_time(1).to_epoch_time() */

    	suite.T().Log("About to run line #35: r.EpochTime(1).ToEpochTime()")

        cursor, err := r.EpochTime(1).ToEpochTime().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #35")
    }

    {
        // times/constructors.yaml line #37
        /* -1 */
        var expected_ int = -1
        /* r.epoch_time(-1).to_epoch_time() */

    	suite.T().Log("About to run line #37: r.EpochTime(-1).ToEpochTime()")

        cursor, err := r.EpochTime(-1).ToEpochTime().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #37")
    }

    {
        // times/constructors.yaml line #39
        /* 1.444 */
        var expected_ float64 = 1.444
        /* r.epoch_time(1.4444445).to_epoch_time() */

    	suite.T().Log("About to run line #39: r.EpochTime(1.4444445).ToEpochTime()")

        cursor, err := r.EpochTime(1.4444445).ToEpochTime().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #39")
    }

    {
        // times/constructors.yaml line #42
        /* "1970-01-01T00:00:01.444+00:00" */
        var expected_ string = "1970-01-01T00:00:01.444+00:00"
        /* r.epoch_time(1.4444445).to_iso8601() */

    	suite.T().Log("About to run line #42: r.EpochTime(1.4444445).ToISO8601()")

        cursor, err := r.EpochTime(1.4444445).ToISO8601().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #42")
    }

    {
        // times/constructors.yaml line #45
        /* 1.444 */
        var expected_ float64 = 1.444
        /* r.epoch_time(1.4444445).seconds() */

    	suite.T().Log("About to run line #45: r.EpochTime(1.4444445).Seconds()")

        cursor, err := r.EpochTime(1.4444445).Seconds().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #45")
    }

    {
        // times/constructors.yaml line #48
        /* 10000 */
        var expected_ int = 10000
        /* r.epoch_time(253430000000).year() */

    	suite.T().Log("About to run line #48: r.EpochTime(253430000000).Year()")

        cursor, err := r.EpochTime(253430000000).Year().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #48")
    }

    {
        // times/constructors.yaml line #50
        /* err("ReqlQueryLogicError", "Year `10000` out of valid ISO 8601 range [0, 9999].", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Year `10000` out of valid ISO 8601 range [0, 9999].")
        /* r.epoch_time(253430000000).to_iso8601() */

    	suite.T().Log("About to run line #50: r.EpochTime(253430000000).ToISO8601()")

        cursor, err := r.EpochTime(253430000000).ToISO8601().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #50")
    }

    {
        // times/constructors.yaml line #53
        /* err("ReqlQueryLogicError", "Error in time logic: Year is out of valid range: 1400..10000.", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Error in time logic: Year is out of valid range: 1400..10000.")
        /* r.epoch_time(253440000000).year() */

    	suite.T().Log("About to run line #53: r.EpochTime(253440000000).Year()")

        cursor, err := r.EpochTime(253440000000).Year().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #53")
    }

    {
        // times/constructors.yaml line #55
        /* 253440000000 */
        var expected_ int = 253440000000
        /* r.epoch_time(253440000000).to_epoch_time() */

    	suite.T().Log("About to run line #55: r.EpochTime(253440000000).ToEpochTime()")

        cursor, err := r.EpochTime(253440000000).ToEpochTime().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #55")
    }

    {
        // times/constructors.yaml line #57
        /* 1400 */
        var expected_ int = 1400
        /* r.epoch_time(-17980000000).year() */

    	suite.T().Log("About to run line #57: r.EpochTime(-17980000000).Year()")

        cursor, err := r.EpochTime(-17980000000).Year().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #57")
    }

    {
        // times/constructors.yaml line #59
        /* err("ReqlQueryLogicError", "Error in time logic: Year is out of valid range: 1400..10000.", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Error in time logic: Year is out of valid range: 1400..10000.")
        /* r.epoch_time(-17990000000).year() */

    	suite.T().Log("About to run line #59: r.EpochTime(-17990000000).Year()")

        cursor, err := r.EpochTime(-17990000000).Year().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #59")
    }

    {
        // times/constructors.yaml line #61
        /* -17990000000 */
        var expected_ int = -17990000000
        /* r.epoch_time(-17990000000).to_epoch_time() */

    	suite.T().Log("About to run line #61: r.EpochTime(-17990000000).ToEpochTime()")

        cursor, err := r.EpochTime(-17990000000).ToEpochTime().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #61")
    }
}
