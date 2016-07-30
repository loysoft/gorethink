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

// Tests RQL json parsing
func TestJsonSuite(t *testing.T) {
	suite.Run(t, new(JsonSuite ))
}

type JsonSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *JsonSuite) SetupTest() {
	suite.T().Log("Setting up JsonSuite")
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

func (suite *JsonSuite) TearDownSuite() {
	suite.T().Log("Tearing down JsonSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *JsonSuite) TestCases() {
	suite.T().Log("Running JsonSuite: Tests RQL json parsing")



	{
		// json.yaml line #4
		/* [1,2,3] */
		var expected_ []interface{} = []interface{}{1, 2, 3}
		/* r.json("[1,2,3]") */

		suite.T().Log("About to run line #4: r.JSON('[1,2,3]')")

		runAndAssert(suite.Suite, expected_, r.JSON("[1,2,3]"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #4")
	}

	{
		// json.yaml line #7
		/* 1 */
		var expected_ int = 1
		/* r.json("1") */

		suite.T().Log("About to run line #7: r.JSON('1')")

		runAndAssert(suite.Suite, expected_, r.JSON("1"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// json.yaml line #10
		/* {} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{}
		/* r.json("{}") */

		suite.T().Log("About to run line #10: r.JSON('{}')")

		runAndAssert(suite.Suite, expected_, r.JSON("{}"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// json.yaml line #13
		/* "foo" */
		var expected_ string = "foo"
		/* r.json('"foo"') */

		suite.T().Log("About to run line #13: r.JSON('\\'foo\\'')")

		runAndAssert(suite.Suite, expected_, r.JSON("\"foo\""), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #13")
	}

	{
		// json.yaml line #16
		/* err("ReqlQueryLogicError", 'Failed to parse "[1,2" as JSON:' + ' Missing a comma or \']\' after an array element.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Failed to parse \"[1,2\" as JSON:" + " Missing a comma or ']' after an array element.")
		/* r.json("[1,2") */

		suite.T().Log("About to run line #16: r.JSON('[1,2')")

		runAndAssert(suite.Suite, expected_, r.JSON("[1,2"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #16")
	}

	{
		// json.yaml line #19
		/* '[1,2,3]' */
		var expected_ string = "[1,2,3]"
		/* r.json("[1,2,3]").to_json_string() */

		suite.T().Log("About to run line #19: r.JSON('[1,2,3]').ToJSON()")

		runAndAssert(suite.Suite, expected_, r.JSON("[1,2,3]").ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #19")
	}

	{
		// json.yaml line #23
		/* '[1,2,3]' */
		var expected_ string = "[1,2,3]"
		/* r.json("[1,2,3]").to_json() */

		suite.T().Log("About to run line #23: r.JSON('[1,2,3]').ToJSON()")

		runAndAssert(suite.Suite, expected_, r.JSON("[1,2,3]").ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// json.yaml line #26
		/* '{"foo":4}' */
		var expected_ string = "{\"foo\":4}"
		/* r.json("{\"foo\":4}").to_json_string() */

		suite.T().Log("About to run line #26: r.JSON('{\\'foo\\':4}').ToJSON()")

		runAndAssert(suite.Suite, expected_, r.JSON("{\"foo\":4}").ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #26")
	}

	{
		// json.yaml line #30
		/* '{"foo":4}' */
		var expected_ string = "{\"foo\":4}"
		/* r.json("{\"foo\":4}").to_json() */

		suite.T().Log("About to run line #30: r.JSON('{\\'foo\\':4}').ToJSON()")

		runAndAssert(suite.Suite, expected_, r.JSON("{\"foo\":4}").ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #30")
	}

	// json.yaml line #34
	// text = '[{"id":1,"first_name":"Harry","last_name":"Riley","email":"hriley0@usgs.gov","country":"Andorra","ip_address":"221.25.65.136"},{"id":2,"first_name":"Bonnie","last_name":"Anderson","email":"banderson1@list-manage.com","country":"Tuvalu","ip_address":"116.162.43.150"},{"id":3,"first_name":"Marie","last_name":"Schmidt","email":"mschmidt2@diigo.com","country":"Iraq","ip_address":"181.105.59.57"},{"id":4,"first_name":"Phillip","last_name":"Willis","email":"pwillis3@com.com","country":"Montenegro","ip_address":"24.223.139.156"}]'
	suite.T().Log("Possibly executing: var text string = '[{\\'id\\':1,\\'first_name\\':\\'Harry\\',\\'last_name\\':\\'Riley\\',\\'email\\':\\'hriley0@usgs.gov\\',\\'country\\':\\'Andorra\\',\\'ip_address\\':\\'221.25.65.136\\'},{\\'id\\':2,\\'first_name\\':\\'Bonnie\\',\\'last_name\\':\\'Anderson\\',\\'email\\':\\'banderson1@list-manage.com\\',\\'country\\':\\'Tuvalu\\',\\'ip_address\\':\\'116.162.43.150\\'},{\\'id\\':3,\\'first_name\\':\\'Marie\\',\\'last_name\\':\\'Schmidt\\',\\'email\\':\\'mschmidt2@diigo.com\\',\\'country\\':\\'Iraq\\',\\'ip_address\\':\\'181.105.59.57\\'},{\\'id\\':4,\\'first_name\\':\\'Phillip\\',\\'last_name\\':\\'Willis\\',\\'email\\':\\'pwillis3@com.com\\',\\'country\\':\\'Montenegro\\',\\'ip_address\\':\\'24.223.139.156\\'}]'")

	text := "[{\"id\":1,\"first_name\":\"Harry\",\"last_name\":\"Riley\",\"email\":\"hriley0@usgs.gov\",\"country\":\"Andorra\",\"ip_address\":\"221.25.65.136\"},{\"id\":2,\"first_name\":\"Bonnie\",\"last_name\":\"Anderson\",\"email\":\"banderson1@list-manage.com\",\"country\":\"Tuvalu\",\"ip_address\":\"116.162.43.150\"},{\"id\":3,\"first_name\":\"Marie\",\"last_name\":\"Schmidt\",\"email\":\"mschmidt2@diigo.com\",\"country\":\"Iraq\",\"ip_address\":\"181.105.59.57\"},{\"id\":4,\"first_name\":\"Phillip\",\"last_name\":\"Willis\",\"email\":\"pwillis3@com.com\",\"country\":\"Montenegro\",\"ip_address\":\"24.223.139.156\"}]"
	_ = text // Prevent any noused variable errors


	// json.yaml line #35
	// sorted = '[{"country":"Andorra","email":"hriley0@usgs.gov","first_name":"Harry","id":1,"ip_address":"221.25.65.136","last_name":"Riley"},{"country":"Tuvalu","email":"banderson1@list-manage.com","first_name":"Bonnie","id":2,"ip_address":"116.162.43.150","last_name":"Anderson"},{"country":"Iraq","email":"mschmidt2@diigo.com","first_name":"Marie","id":3,"ip_address":"181.105.59.57","last_name":"Schmidt"},{"country":"Montenegro","email":"pwillis3@com.com","first_name":"Phillip","id":4,"ip_address":"24.223.139.156","last_name":"Willis"}]'
	suite.T().Log("Possibly executing: var sorted string = '[{\\'country\\':\\'Andorra\\',\\'email\\':\\'hriley0@usgs.gov\\',\\'first_name\\':\\'Harry\\',\\'id\\':1,\\'ip_address\\':\\'221.25.65.136\\',\\'last_name\\':\\'Riley\\'},{\\'country\\':\\'Tuvalu\\',\\'email\\':\\'banderson1@list-manage.com\\',\\'first_name\\':\\'Bonnie\\',\\'id\\':2,\\'ip_address\\':\\'116.162.43.150\\',\\'last_name\\':\\'Anderson\\'},{\\'country\\':\\'Iraq\\',\\'email\\':\\'mschmidt2@diigo.com\\',\\'first_name\\':\\'Marie\\',\\'id\\':3,\\'ip_address\\':\\'181.105.59.57\\',\\'last_name\\':\\'Schmidt\\'},{\\'country\\':\\'Montenegro\\',\\'email\\':\\'pwillis3@com.com\\',\\'first_name\\':\\'Phillip\\',\\'id\\':4,\\'ip_address\\':\\'24.223.139.156\\',\\'last_name\\':\\'Willis\\'}]'")

	sorted := "[{\"country\":\"Andorra\",\"email\":\"hriley0@usgs.gov\",\"first_name\":\"Harry\",\"id\":1,\"ip_address\":\"221.25.65.136\",\"last_name\":\"Riley\"},{\"country\":\"Tuvalu\",\"email\":\"banderson1@list-manage.com\",\"first_name\":\"Bonnie\",\"id\":2,\"ip_address\":\"116.162.43.150\",\"last_name\":\"Anderson\"},{\"country\":\"Iraq\",\"email\":\"mschmidt2@diigo.com\",\"first_name\":\"Marie\",\"id\":3,\"ip_address\":\"181.105.59.57\",\"last_name\":\"Schmidt\"},{\"country\":\"Montenegro\",\"email\":\"pwillis3@com.com\",\"first_name\":\"Phillip\",\"id\":4,\"ip_address\":\"24.223.139.156\",\"last_name\":\"Willis\"}]"
	_ = sorted // Prevent any noused variable errors


	{
		// json.yaml line #37
		/* sorted */
		var expected_ string = sorted
		/* r.json(text).to_json_string() */

		suite.T().Log("About to run line #37: r.JSON(text).ToJSON()")

		runAndAssert(suite.Suite, expected_, r.JSON(text).ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// json.yaml line #40
		/* err('ReqlQueryLogicError', 'Cannot convert `r.minval` to JSON.') */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot convert `r.minval` to JSON.")
		/* r.expr(r.minval).to_json_string() */

		suite.T().Log("About to run line #40: r.Expr(r.MinVal).ToJSON()")

		runAndAssert(suite.Suite, expected_, r.Expr(r.MinVal).ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #40")
	}

	{
		// json.yaml line #43
		/* err('ReqlQueryLogicError', 'Cannot convert `r.maxval` to JSON.') */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot convert `r.maxval` to JSON.")
		/* r.expr(r.maxval).to_json_string() */

		suite.T().Log("About to run line #43: r.Expr(r.MaxVal).ToJSON()")

		runAndAssert(suite.Suite, expected_, r.Expr(r.MaxVal).ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #43")
	}

	{
		// json.yaml line #46
		/* err('ReqlQueryLogicError', 'Cannot convert `r.minval` to JSON.') */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot convert `r.minval` to JSON.")
		/* r.expr(r.minval).coerce_to('string') */

		suite.T().Log("About to run line #46: r.Expr(r.MinVal).CoerceTo('string')")

		runAndAssert(suite.Suite, expected_, r.Expr(r.MinVal).CoerceTo("string"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #46")
	}

	{
		// json.yaml line #49
		/* err('ReqlQueryLogicError', 'Cannot convert `r.maxval` to JSON.') */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot convert `r.maxval` to JSON.")
		/* r.expr(r.maxval).coerce_to('string') */

		suite.T().Log("About to run line #49: r.Expr(r.MaxVal).CoerceTo('string')")

		runAndAssert(suite.Suite, expected_, r.Expr(r.MaxVal).CoerceTo("string"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #49")
	}

	{
		// json.yaml line #52
		/* {'timezone':'+00:00','$reql_type$':'TIME','epoch_time':1410393600} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"timezone": "+00:00", "$reql_type$": "TIME", "epoch_time": 1410393600, }
		/* r.time(2014,9,11, 'Z') */

		suite.T().Log("About to run line #52: r.Time(2014, 9, 11, 'Z')")

		runAndAssert(suite.Suite, expected_, r.Time(2014, 9, 11, "Z"), suite.session, r.RunOpts{
			TimeFormat: "raw",
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #52")
	}

	{
		// json.yaml line #57
		/* '{"$reql_type$":"TIME","epoch_time":1410393600,"timezone":"+00:00"}' */
		var expected_ string = "{\"$reql_type$\":\"TIME\",\"epoch_time\":1410393600,\"timezone\":\"+00:00\"}"
		/* r.time(2014,9,11, 'Z').to_json_string() */

		suite.T().Log("About to run line #57: r.Time(2014, 9, 11, 'Z').ToJSON()")

		runAndAssert(suite.Suite, expected_, r.Time(2014, 9, 11, "Z").ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #57")
	}

	{
		// json.yaml line #60
		/* {'$reql_type$':'GEOMETRY','coordinates':[0,0],'type':'Point'} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{0, 0}, "type": "Point", }
		/* r.point(0,0) */

		suite.T().Log("About to run line #60: r.Point(0, 0)")

		runAndAssert(suite.Suite, expected_, r.Point(0, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #60")
	}

	{
		// json.yaml line #63
		/* '{"$reql_type$":"GEOMETRY","coordinates":[0,0],"type":"Point"}' */
		var expected_ string = "{\"$reql_type$\":\"GEOMETRY\",\"coordinates\":[0,0],\"type\":\"Point\"}"
		/* r.point(0,0).to_json_string() */

		suite.T().Log("About to run line #63: r.Point(0, 0).ToJSON()")

		runAndAssert(suite.Suite, expected_, r.Point(0, 0).ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #63")
	}

	// json.yaml line #68
	// s = b'\x66\x6f\x6f'
	suite.T().Log("Possibly executing: var s []byte = []byte{102,111,111}")

	s := []byte{102,111,111}
	_ = s // Prevent any noused variable errors


	{
		// json.yaml line #70
		/* s */
		var expected_ []byte = s
		/* r.binary(s) */

		suite.T().Log("About to run line #70: r.Binary(s)")

		runAndAssert(suite.Suite, expected_, r.Binary(s), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #70")
	}

	{
		// json.yaml line #73
		/* '{"$reql_type$":"BINARY","data":"Zm9v"}' */
		var expected_ string = "{\"$reql_type$\":\"BINARY\",\"data\":\"Zm9v\"}"
		/* r.expr("foo").coerce_to("binary").to_json_string() */

		suite.T().Log("About to run line #73: r.Expr('foo').CoerceTo('binary').ToJSON()")

		runAndAssert(suite.Suite, expected_, r.Expr("foo").CoerceTo("binary").ToJSON(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #73")
	}
}
