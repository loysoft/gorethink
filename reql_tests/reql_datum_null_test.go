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
func TestDatumNullSuite(t *testing.T) {
    suite.Run(t, new(DatumNullSuite ))
}

type DatumNullSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *DatumNullSuite) SetupTest() {
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

func (suite *DatumNullSuite) TearDownSuite() {
	r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
    r.DBDrop("test").Exec(suite.session)

    suite.session.Close()
}

func (suite *DatumNullSuite) TestCases() {


    {
        // datum/null.yaml line #6
        /* (null) */
        var expected_ interface{} = nil
        /* r.expr(null) */

    	suite.T().Log("About to run line #6: r.Expr(nil)")

        cursor, err := r.Expr(nil).Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #6")
    }

    {
        // datum/null.yaml line #9
        /* 'NULL' */
        var expected_ string = "NULL"
        /* r.expr(null).type_of() */

    	suite.T().Log("About to run line #9: r.Expr(nil).TypeOf()")

        cursor, err := r.Expr(nil).TypeOf().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #9")
    }

    {
        // datum/null.yaml line #14
        /* 'null' */
        var expected_ string = "null"
        /* r.expr(null).coerce_to('string') */

    	suite.T().Log("About to run line #14: r.Expr(nil).CoerceTo('string')")

        cursor, err := r.Expr(nil).CoerceTo("string").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #14")
    }

    {
        // datum/null.yaml line #17
        /* null */
        var expected_ interface{} = nil
        /* r.expr(null).coerce_to('null') */

    	suite.T().Log("About to run line #17: r.Expr(nil).CoerceTo('null')")

        cursor, err := r.Expr(nil).CoerceTo("null").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #17")
    }
}
