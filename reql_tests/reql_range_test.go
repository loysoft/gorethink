// Autogenerated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../../gen_tests/template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/dancannon/gorethink.v2"
)

// Tests RQL range generation
func TestRangeSuite(t *testing.T) {
    suite.Run(t, new(RangeSuite ))
}

type RangeSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *RangeSuite) SetupTest() {
	suite.T().Log("Setting up RangeSuite")
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

func (suite *RangeSuite) TearDownSuite() {
	suite.T().Log("Tearing down RangeSuite")

	r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
    r.DBDrop("test").Exec(suite.session)

    suite.session.Close()
}

func (suite *RangeSuite) TestCases() {
	suite.T().Log("Running RangeSuite: Tests RQL range generation")



    {
        // range.yaml line #3
        /* 'STREAM' */
        var expected_ string = "STREAM"
        /* r.range().type_of() */

    	suite.T().Log("About to run line #3: r.Range().TypeOf()")

        runAndAssert(suite.Suite, expected_, r.Range().TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #3")
    }

    {
        // range.yaml line #6
        /* [0, 1, 2, 3] */
        var expected_ []interface{} = []interface{}{0, 1, 2, 3}
        /* r.range().limit(4) */

    	suite.T().Log("About to run line #6: r.Range().Limit(4)")

        runAndAssert(suite.Suite, expected_, r.Range().Limit(4), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #6")
    }

    {
        // range.yaml line #9
        /* [0, 1, 2, 3] */
        var expected_ []interface{} = []interface{}{0, 1, 2, 3}
        /* r.range(4) */

    	suite.T().Log("About to run line #9: r.Range(4)")

        runAndAssert(suite.Suite, expected_, r.Range(4), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #9")
    }

    {
        // range.yaml line #12
        /* [2, 3, 4] */
        var expected_ []interface{} = []interface{}{2, 3, 4}
        /* r.range(2, 5) */

    	suite.T().Log("About to run line #12: r.Range(2, 5)")

        runAndAssert(suite.Suite, expected_, r.Range(2, 5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #12")
    }

    {
        // range.yaml line #15
        /* [] */
        var expected_ []interface{} = []interface{}{}
        /* r.range(0) */

    	suite.T().Log("About to run line #15: r.Range(0)")

        runAndAssert(suite.Suite, expected_, r.Range(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #15")
    }

    {
        // range.yaml line #18
        /* [] */
        var expected_ []interface{} = []interface{}{}
        /* r.range(5, 2) */

    	suite.T().Log("About to run line #18: r.Range(5, 2)")

        runAndAssert(suite.Suite, expected_, r.Range(5, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #18")
    }

    {
        // range.yaml line #21
        /* [-5, -4, -3] */
        var expected_ []interface{} = []interface{}{-5, -4, -3}
        /* r.range(-5, -2) */

    	suite.T().Log("About to run line #21: r.Range(-5, -2)")

        runAndAssert(suite.Suite, expected_, r.Range(-5, -2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #21")
    }

    {
        // range.yaml line #24
        /* [-5, -4, -3, -2, -1, 0, 1] */
        var expected_ []interface{} = []interface{}{-5, -4, -3, -2, -1, 0, 1}
        /* r.range(-5, 2) */

    	suite.T().Log("About to run line #24: r.Range(-5, 2)")

        runAndAssert(suite.Suite, expected_, r.Range(-5, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #24")
    }

    {
        // range.yaml line #27
        /* err("ReqlCompileError", "Expected between 0 and 2 arguments but found 3.", []) */
        var expected_ Err = err("ReqlCompileError", "Expected between 0 and 2 arguments but found 3.")
        /* r.range(2, 5, 8) */

    	suite.T().Log("About to run line #27: r.Range(2, 5, 8)")

        runAndAssert(suite.Suite, expected_, r.Range(2, 5, 8), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #27")
    }

    {
        // range.yaml line #30
        /* err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
        /* r.range("foo") */

    	suite.T().Log("About to run line #30: r.Range('foo')")

        runAndAssert(suite.Suite, expected_, r.Range("foo"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #30")
    }

    {
        // range.yaml line #34
        /* err_regex("ReqlQueryLogicError", "Number not an integer \\(>2\\^53\\). 9007199254740994", []) */
        var expected_ ErrRegex = err_regex("ReqlQueryLogicError", "Number not an integer \\(>2\\^53\\). 9007199254740994")
        /* r.range(9007199254740994) */

    	suite.T().Log("About to run line #34: r.Range(9007199254740994)")

        runAndAssert(suite.Suite, expected_, r.Range(9007199254740994), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #34")
    }

    {
        // range.yaml line #37
        /* err_regex("ReqlQueryLogicError", "Number not an integer \\(<-2\\^53\\). -9007199254740994", []) */
        var expected_ ErrRegex = err_regex("ReqlQueryLogicError", "Number not an integer \\(<-2\\^53\\). -9007199254740994")
        /* r.range(-9007199254740994) */

    	suite.T().Log("About to run line #37: r.Range(-9007199254740994)")

        runAndAssert(suite.Suite, expected_, r.Range(-9007199254740994), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #37")
    }

    {
        // range.yaml line #40
        /* err_regex("ReqlQueryLogicError", "Number not an integer. 0\\.5", []) */
        var expected_ ErrRegex = err_regex("ReqlQueryLogicError", "Number not an integer. 0\\.5")
        /* r.range(0.5) */

    	suite.T().Log("About to run line #40: r.Range(0.5)")

        runAndAssert(suite.Suite, expected_, r.Range(0.5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #40")
    }

    {
        // range.yaml line #43
        /* err("ReqlQueryLogicError", "Cannot use an infinite stream with an aggregation function (`reduce`, `count`, etc.) or coerce it to an array.", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Cannot use an infinite stream with an aggregation function (`reduce`, `count`, etc.) or coerce it to an array.")
        /* r.range().count() */

    	suite.T().Log("About to run line #43: r.Range().Count()")

        runAndAssert(suite.Suite, expected_, r.Range().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #43")
    }

    {
        // range.yaml line #46
        /* err("ReqlQueryLogicError", "Cannot use an infinite stream with an aggregation function (`reduce`, `count`, etc.) or coerce it to an array.", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Cannot use an infinite stream with an aggregation function (`reduce`, `count`, etc.) or coerce it to an array.")
        /* r.range().coerce_to("ARRAY") */

    	suite.T().Log("About to run line #46: r.Range().CoerceTo('ARRAY')")

        runAndAssert(suite.Suite, expected_, r.Range().CoerceTo("ARRAY"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #46")
    }

    {
        // range.yaml line #49
        /* err("ReqlQueryLogicError", "Cannot use an infinite stream with an aggregation function (`reduce`, `count`, etc.) or coerce it to an array.", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Cannot use an infinite stream with an aggregation function (`reduce`, `count`, etc.) or coerce it to an array.")
        /* r.range().coerce_to("OBJECT") */

    	suite.T().Log("About to run line #49: r.Range().CoerceTo('OBJECT')")

        runAndAssert(suite.Suite, expected_, r.Range().CoerceTo("OBJECT"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #49")
    }

    {
        // range.yaml line #52
        /* 4 */
        var expected_ int = 4
        /* r.range(4).count() */

    	suite.T().Log("About to run line #52: r.Range(4).Count()")

        runAndAssert(suite.Suite, expected_, r.Range(4).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #52")
    }
}
