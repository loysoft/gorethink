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
func TestMetaDbsSuite(t *testing.T) {
    suite.Run(t, new(MetaDbsSuite ))
}

type MetaDbsSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MetaDbsSuite) SetupTest() {
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

func (suite *MetaDbsSuite) TearDownSuite() {
	r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
    r.DBDrop("test").Exec(suite.session)

    suite.session.Close()
}

func (suite *MetaDbsSuite) TestCases() {


    {
        // meta/dbs.yaml line #6
        /* bag(['rethinkdb', 'test']) */
        var expected_ Expected = bag([]interface{}{"rethinkdb", "test"})
        /* r.db_list() */

    	suite.T().Log("About to run line #6: r.DBList()")

        cursor, err := r.DBList().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #6")
    }

    {
        // meta/dbs.yaml line #11
        /* partial({'dbs_created':1}) */
        var expected_ Expected = partial(map[interface{}]interface{}{"dbs_created": 1, })
        /* r.db_create('a') */

    	suite.T().Log("About to run line #11: r.DBCreate('a')")

        cursor, err := r.DBCreate("a").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #11")
    }

    {
        // meta/dbs.yaml line #13
        /* partial({'dbs_created':1}) */
        var expected_ Expected = partial(map[interface{}]interface{}{"dbs_created": 1, })
        /* r.db_create('b') */

    	suite.T().Log("About to run line #13: r.DBCreate('b')")

        cursor, err := r.DBCreate("b").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #13")
    }

    {
        // meta/dbs.yaml line #18
        /* bag(['rethinkdb', 'a', 'b', 'test']) */
        var expected_ Expected = bag([]interface{}{"rethinkdb", "a", "b", "test"})
        /* r.db_list() */

    	suite.T().Log("About to run line #18: r.DBList()")

        cursor, err := r.DBList().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #18")
    }

    {
        // meta/dbs.yaml line #23
        /* {'name':'a','id':uuid()} */
        var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"name": "a", "id": uuid(), }
        /* r.db('a').config() */

    	suite.T().Log("About to run line #23: r.DB('a').Config()")

        cursor, err := r.DB("a").Config().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #23")
    }

    {
        // meta/dbs.yaml line #28
        /* partial({'dbs_dropped':1}) */
        var expected_ Expected = partial(map[interface{}]interface{}{"dbs_dropped": 1, })
        /* r.db_drop('b') */

    	suite.T().Log("About to run line #28: r.DBDrop('b')")

        cursor, err := r.DBDrop("b").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #28")
    }

    {
        // meta/dbs.yaml line #31
        /* bag(['rethinkdb', 'a', 'test']) */
        var expected_ Expected = bag([]interface{}{"rethinkdb", "a", "test"})
        /* r.db_list() */

    	suite.T().Log("About to run line #31: r.DBList()")

        cursor, err := r.DBList().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #31")
    }

    {
        // meta/dbs.yaml line #34
        /* partial({'dbs_dropped':1}) */
        var expected_ Expected = partial(map[interface{}]interface{}{"dbs_dropped": 1, })
        /* r.db_drop('a') */

    	suite.T().Log("About to run line #34: r.DBDrop('a')")

        cursor, err := r.DBDrop("a").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #34")
    }

    {
        // meta/dbs.yaml line #37
        /* bag(['rethinkdb', 'test']) */
        var expected_ Expected = bag([]interface{}{"rethinkdb", "test"})
        /* r.db_list() */

    	suite.T().Log("About to run line #37: r.DBList()")

        cursor, err := r.DBList().Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #37")
    }

    {
        // meta/dbs.yaml line #41
        /* partial({'dbs_created':1}) */
        var expected_ Expected = partial(map[interface{}]interface{}{"dbs_created": 1, })
        /* r.db_create('bar') */

    	suite.T().Log("About to run line #41: r.DBCreate('bar')")

        cursor, err := r.DBCreate("bar").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #41")
    }

    {
        // meta/dbs.yaml line #44
        /* err('ReqlOpFailedError', 'Database `bar` already exists.', [0]) */
        var expected_ Err = err("ReqlOpFailedError", "Database `bar` already exists.")
        /* r.db_create('bar') */

    	suite.T().Log("About to run line #44: r.DBCreate('bar')")

        cursor, err := r.DBCreate("bar").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #44")
    }

    {
        // meta/dbs.yaml line #47
        /* partial({'dbs_dropped':1}) */
        var expected_ Expected = partial(map[interface{}]interface{}{"dbs_dropped": 1, })
        /* r.db_drop('bar') */

    	suite.T().Log("About to run line #47: r.DBDrop('bar')")

        cursor, err := r.DBDrop("bar").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #47")
    }

    {
        // meta/dbs.yaml line #50
        /* err('ReqlOpFailedError', 'Database `bar` does not exist.', [0]) */
        var expected_ Err = err("ReqlOpFailedError", "Database `bar` does not exist.")
        /* r.db_drop('bar') */

    	suite.T().Log("About to run line #50: r.DBDrop('bar')")

        cursor, err := r.DBDrop("bar").Run(suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
        suite.T().Log("Finished running line #50")
    }
}
