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
func TestTimesShimSuite(t *testing.T) {
    suite.Run(t, new(TimesShimSuite ))
}

type TimesShimSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TimesShimSuite) SetupTest() {
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

func (suite *TimesShimSuite) TearDownSuite() {
	r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
    r.DBDrop("test").Exec(suite.session)

    suite.session.Close()
}

func (suite *TimesShimSuite) TestCases() {


    // times/shim.yaml line #4
    // t = 1375147296.68
    suite.T().Log("Possibly executing: var t float64 = 1375147296.68")

    var t float64 = 1375147296.68
	

    {
        // times/shim.yaml line #8
        /* ("2013-07-29T18:21:36.680-07:00") */
        var expected_ string = "2013-07-29T18:21:36.680-07:00"
        /* r.expr(datetime.fromtimestamp(t, PacificTimeZone())).to_iso8601() */

    	suite.T().Log("About to run line #8: r.Expr(Ast.Fromtimestamp(t, PacificTimeZone())).ToISO8601()")

        cursor, err := r.Expr(Ast.Fromtimestamp(t, PacificTimeZone())).ToISO8601().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #8")
    }

    {
        // times/shim.yaml line #12
        /* ("2013-07-30T01:21:36.680+00:00") */
        var expected_ string = "2013-07-30T01:21:36.680+00:00"
        /* r.expr(datetime.fromtimestamp(t, UTCTimeZone())).to_iso8601() */

    	suite.T().Log("About to run line #12: r.Expr(Ast.Fromtimestamp(t, UTCTimeZone())).ToISO8601()")

        cursor, err := r.Expr(Ast.Fromtimestamp(t, UTCTimeZone())).ToISO8601().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #12")
    }

    {
        // times/shim.yaml line #16
        /* (1375147296.68) */
        var expected_ float64 = 1375147296.68
        /* r.expr(datetime.fromtimestamp(t, PacificTimeZone())).to_epoch_time() */

    	suite.T().Log("About to run line #16: r.Expr(Ast.Fromtimestamp(t, PacificTimeZone())).ToEpochTime()")

        cursor, err := r.Expr(Ast.Fromtimestamp(t, PacificTimeZone())).ToEpochTime().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #16")
    }

    {
        // times/shim.yaml line #20
        /* (1375147296.68) */
        var expected_ float64 = 1375147296.68
        /* r.expr(datetime.fromtimestamp(t, UTCTimeZone())).to_epoch_time() */

    	suite.T().Log("About to run line #20: r.Expr(Ast.Fromtimestamp(t, UTCTimeZone())).ToEpochTime()")

        cursor, err := r.Expr(Ast.Fromtimestamp(t, UTCTimeZone())).ToEpochTime().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #20")
    }
}
