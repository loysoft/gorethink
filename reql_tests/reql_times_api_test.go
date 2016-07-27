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

// date/time api (#977)
func TestTimesApiSuite(t *testing.T) {
    suite.Run(t, new(TimesApiSuite ))
}

type TimesApiSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TimesApiSuite) SetupTest() {
	suite.T().Log("Setting up TimesApiSuite")
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

func (suite *TimesApiSuite) TearDownSuite() {
	suite.T().Log("Tearing down TimesApiSuite")

	r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
    r.DBDrop("test").Exec(suite.session)

    suite.session.Close()
}

func (suite *TimesApiSuite) TestCases() {
	suite.T().Log("Running TimesApiSuite: date/time api (#977)")



    // times/api.yaml line #6
    // rt1 = 1375147296.6812
    suite.T().Log("Possibly executing: var rt1 float64 = 1375147296.6812")

    var rt1 float64 = 1375147296.6812
    _ = rt1 // Prevent any noused variable errors


    // times/api.yaml line #7
    // t1 = r.epoch_time(rt1)
    suite.T().Log("Possibly executing: var t1 r.Term = r.EpochTime(rt1)")

    var t1 r.Term = r.EpochTime(rt1)
    _ = t1 // Prevent any noused variable errors


    // times/api.yaml line #8
    // t2 = r.epoch_time(rt1 + 1000)
    suite.T().Log("Possibly executing: var t2 r.Term = r.EpochTime(r.Add(rt1, 1000))")

    var t2 r.Term = r.EpochTime(r.Add(rt1, 1000))
    _ = t2 // Prevent any noused variable errors


    {
        // times/api.yaml line #11
        /* (1375148296.681) */
        var expected_ float64 = 1375148296.681
        /* (t1 + 1000).to_epoch_time() */

    	suite.T().Log("About to run line #11: r.Add(t1, 1000).ToEpochTime()")

        runAndAssert(suite.Suite, expected_, r.Add(t1, 1000).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #11")
    }

    {
        // times/api.yaml line #14
        /* (1375146296.681) */
        var expected_ float64 = 1375146296.681
        /* (t1 - 1000).to_epoch_time() */

    	suite.T().Log("About to run line #14: r.Sub(t1, 1000).ToEpochTime()")

        runAndAssert(suite.Suite, expected_, r.Sub(t1, 1000).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #14")
    }

    {
        // times/api.yaml line #17
        /* 1000 */
        var expected_ int = 1000
        /* (t1 - (t1 - 1000)) */

    	suite.T().Log("About to run line #17: r.Sub(t1, r.Sub(t1, 1000))")

        runAndAssert(suite.Suite, expected_, r.Sub(t1, r.Sub(t1, 1000)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #17")
    }

    {
        // times/api.yaml line #22
        /* false */
        var expected_ bool = false
        /* (t1 < t1) */

    	suite.T().Log("About to run line #22: r.Lt(t1, t1)")

        runAndAssert(suite.Suite, expected_, r.Lt(t1, t1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #22")
    }

    {
        // times/api.yaml line #25
        /* true */
        var expected_ bool = true
        /* (t1 <= t1) */

    	suite.T().Log("About to run line #25: r.Le(t1, t1)")

        runAndAssert(suite.Suite, expected_, r.Le(t1, t1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #25")
    }

    {
        // times/api.yaml line #29
        /* true */
        var expected_ bool = true
        /* (t1 == t1) */

    	suite.T().Log("About to run line #29: r.Eq(t1, t1)")

        runAndAssert(suite.Suite, expected_, r.Eq(t1, t1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #29")
    }

    {
        // times/api.yaml line #32
        /* false */
        var expected_ bool = false
        /* (t1 != t1) */

    	suite.T().Log("About to run line #32: r.Ne(t1, t1)")

        runAndAssert(suite.Suite, expected_, r.Ne(t1, t1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #32")
    }

    {
        // times/api.yaml line #34
        /* true */
        var expected_ bool = true
        /* (t1 >= t1) */

    	suite.T().Log("About to run line #34: r.Ge(t1, t1)")

        runAndAssert(suite.Suite, expected_, r.Ge(t1, t1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #34")
    }

    {
        // times/api.yaml line #37
        /* false */
        var expected_ bool = false
        /* (t1 > t1) */

    	suite.T().Log("About to run line #37: r.Gt(t1, t1)")

        runAndAssert(suite.Suite, expected_, r.Gt(t1, t1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #37")
    }

    {
        // times/api.yaml line #40
        /* true */
        var expected_ bool = true
        /* (t1 < t2) */

    	suite.T().Log("About to run line #40: r.Lt(t1, t2)")

        runAndAssert(suite.Suite, expected_, r.Lt(t1, t2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #40")
    }

    {
        // times/api.yaml line #43
        /* true */
        var expected_ bool = true
        /* (t1 <= t2) */

    	suite.T().Log("About to run line #43: r.Le(t1, t2)")

        runAndAssert(suite.Suite, expected_, r.Le(t1, t2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #43")
    }

    {
        // times/api.yaml line #47
        /* false */
        var expected_ bool = false
        /* (t1 == t2) */

    	suite.T().Log("About to run line #47: r.Eq(t1, t2)")

        runAndAssert(suite.Suite, expected_, r.Eq(t1, t2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #47")
    }

    {
        // times/api.yaml line #50
        /* true */
        var expected_ bool = true
        /* (t1 != t2) */

    	suite.T().Log("About to run line #50: r.Ne(t1, t2)")

        runAndAssert(suite.Suite, expected_, r.Ne(t1, t2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #50")
    }

    {
        // times/api.yaml line #52
        /* false */
        var expected_ bool = false
        /* (t1 >= t2) */

    	suite.T().Log("About to run line #52: r.Ge(t1, t2)")

        runAndAssert(suite.Suite, expected_, r.Ge(t1, t2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #52")
    }

    {
        // times/api.yaml line #55
        /* false */
        var expected_ bool = false
        /* (t1 > t2) */

    	suite.T().Log("About to run line #55: r.Gt(t1, t2)")

        runAndAssert(suite.Suite, expected_, r.Gt(t1, t2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #55")
    }

    {
        // times/api.yaml line #60
        /* true */
        var expected_ bool = true
        /* t1.during(t1, t1 + 1000) */

    	suite.T().Log("About to run line #60: t1.During(t1, r.Add(t1, 1000))")

        runAndAssert(suite.Suite, expected_, t1.During(t1, r.Add(t1, 1000)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #60")
    }

    {
        // times/api.yaml line #64
        /* false */
        var expected_ bool = false
        /* t1.during(t1, t1 + 1000, left_bound='open') */

    	suite.T().Log("About to run line #64: t1.During(t1, r.Add(t1, 1000), r.DuringOpts{LeftBound: 'open', })")

        runAndAssert(suite.Suite, expected_, t1.During(t1, r.Add(t1, 1000), r.DuringOpts{LeftBound: "open", }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #64")
    }

    {
        // times/api.yaml line #67
        /* false */
        var expected_ bool = false
        /* t1.during(t1, t1) */

    	suite.T().Log("About to run line #67: t1.During(t1, t1)")

        runAndAssert(suite.Suite, expected_, t1.During(t1, t1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #67")
    }

    {
        // times/api.yaml line #70
        /* true */
        var expected_ bool = true
        /* t1.during(t1, t1, right_bound='closed') */

    	suite.T().Log("About to run line #70: t1.During(t1, t1, r.DuringOpts{RightBound: 'closed', })")

        runAndAssert(suite.Suite, expected_, t1.During(t1, t1, r.DuringOpts{RightBound: "closed", }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #70")
    }

    {
        // times/api.yaml line #77
        /* 1375142400 */
        var expected_ int = 1375142400
        /* t1.date().to_epoch_time() */

    	suite.T().Log("About to run line #77: t1.Date().ToEpochTime()")

        runAndAssert(suite.Suite, expected_, t1.Date().ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #77")
    }

    {
        // times/api.yaml line #79
        /* (4896.681) */
        var expected_ float64 = 4896.681
        /* t1.time_of_day() */

    	suite.T().Log("About to run line #79: t1.TimeOfDay()")

        runAndAssert(suite.Suite, expected_, t1.TimeOfDay(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #79")
    }

    {
        // times/api.yaml line #81
        /* 2013 */
        var expected_ int = 2013
        /* t1.year() */

    	suite.T().Log("About to run line #81: t1.Year()")

        runAndAssert(suite.Suite, expected_, t1.Year(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #81")
    }

    {
        // times/api.yaml line #83
        /* 7 */
        var expected_ int = 7
        /* t1.month() */

    	suite.T().Log("About to run line #83: t1.Month()")

        runAndAssert(suite.Suite, expected_, t1.Month(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #83")
    }

    {
        // times/api.yaml line #85
        /* 30 */
        var expected_ int = 30
        /* t1.day() */

    	suite.T().Log("About to run line #85: t1.Day()")

        runAndAssert(suite.Suite, expected_, t1.Day(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #85")
    }

    {
        // times/api.yaml line #87
        /* 2 */
        var expected_ int = 2
        /* t1.day_of_week() */

    	suite.T().Log("About to run line #87: t1.DayOfWeek()")

        runAndAssert(suite.Suite, expected_, t1.DayOfWeek(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #87")
    }

    {
        // times/api.yaml line #89
        /* 211 */
        var expected_ int = 211
        /* t1.day_of_year() */

    	suite.T().Log("About to run line #89: t1.DayOfYear()")

        runAndAssert(suite.Suite, expected_, t1.DayOfYear(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #89")
    }

    {
        // times/api.yaml line #91
        /* 1 */
        var expected_ int = 1
        /* t1.hours() */

    	suite.T().Log("About to run line #91: t1.Hours()")

        runAndAssert(suite.Suite, expected_, t1.Hours(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #91")
    }

    {
        // times/api.yaml line #93
        /* 21 */
        var expected_ int = 21
        /* t1.minutes() */

    	suite.T().Log("About to run line #93: t1.Minutes()")

        runAndAssert(suite.Suite, expected_, t1.Minutes(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #93")
    }

    {
        // times/api.yaml line #95
        /* 36.681 */
        var expected_ float64 = 36.681
        /* t1.seconds() */

    	suite.T().Log("About to run line #95: t1.Seconds()")

        runAndAssert(suite.Suite, expected_, t1.Seconds(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #95")
    }

    {
        // times/api.yaml line #99
        /* (1375165800.1) */
        var expected_ float64 = 1375165800.1
        /* r.time(2013, r.july, 29, 23, 30, 0.1, "-07:00").to_epoch_time() */

    	suite.T().Log("About to run line #99: r.Time(2013, r.July, 29, 23, 30, 0.1, '-07:00').ToEpochTime()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29, 23, 30, 0.1, "-07:00").ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #99")
    }

    {
        // times/api.yaml line #101
        /* ("-07:00") */
        var expected_ string = "-07:00"
        /* r.time(2013, r.july, 29, 23, 30, 0.1, "-07:00").timezone() */

    	suite.T().Log("About to run line #101: r.Time(2013, r.July, 29, 23, 30, 0.1, '-07:00').Timezone()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29, 23, 30, 0.1, "-07:00").Timezone(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #101")
    }

    {
        // times/api.yaml line #103
        /* err("ReqlQueryLogicError", "Got 6 arguments to TIME (expected 4 or 7).", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Got 6 arguments to TIME (expected 4 or 7).")
        /* r.time(2013, r.july, 29, 23, 30, 0.1).to_epoch_time() */

    	suite.T().Log("About to run line #103: r.Time(2013, r.July, 29, 23, 30, 0.1).ToEpochTime()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29, 23, 30, 0.1).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #103")
    }

    {
        // times/api.yaml line #105
        /* err("ReqlQueryLogicError", "Got 6 arguments to TIME (expected 4 or 7).", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Got 6 arguments to TIME (expected 4 or 7).")
        /* r.time(2013, r.july, 29, 23, 30, 0.1).timezone() */

    	suite.T().Log("About to run line #105: r.Time(2013, r.July, 29, 23, 30, 0.1).Timezone()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29, 23, 30, 0.1).Timezone(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #105")
    }

    {
        // times/api.yaml line #107
        /* err("ReqlQueryLogicError", "Got 5 arguments to TIME (expected 4 or 7).", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Got 5 arguments to TIME (expected 4 or 7).")
        /* r.time(2013, r.july, 29, 23, 30).to_epoch_time() */

    	suite.T().Log("About to run line #107: r.Time(2013, r.July, 29, 23, 30).ToEpochTime()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29, 23, 30).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #107")
    }

    {
        // times/api.yaml line #109
        /* err("ReqlQueryLogicError", "Expected type STRING but found NUMBER.", []) */
        var expected_ Err = err("ReqlQueryLogicError", "Expected type STRING but found NUMBER.")
        /* r.time(2013, r.july, 29, 23).to_epoch_time() */

    	suite.T().Log("About to run line #109: r.Time(2013, r.July, 29, 23).ToEpochTime()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29, 23).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #109")
    }

    {
        // times/api.yaml line #111
        /* 1375081200 */
        var expected_ int = 1375081200
        /* r.time(2013, r.july, 29, "-07:00").to_epoch_time() */

    	suite.T().Log("About to run line #111: r.Time(2013, r.July, 29, '-07:00').ToEpochTime()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29, "-07:00").ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #111")
    }

    {
        // times/api.yaml line #113
        /* ("-07:00") */
        var expected_ string = "-07:00"
        /* r.time(2013, r.july, 29, "-07:00").timezone() */

    	suite.T().Log("About to run line #113: r.Time(2013, r.July, 29, '-07:00').Timezone()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29, "-07:00").Timezone(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #113")
    }

    {
        // times/api.yaml line #115
        /* err("ReqlCompileError", "Expected between 4 and 7 arguments but found 3.", []) */
        var expected_ Err = err("ReqlCompileError", "Expected between 4 and 7 arguments but found 3.")
        /* r.time(2013, r.july, 29).to_epoch_time() */

    	suite.T().Log("About to run line #115: r.Time(2013, r.July, 29).ToEpochTime()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #115")
    }

    {
        // times/api.yaml line #117
        /* err("ReqlCompileError", "Expected between 4 and 7 arguments but found 3.", []) */
        var expected_ Err = err("ReqlCompileError", "Expected between 4 and 7 arguments but found 3.")
        /* r.time(2013, r.july, 29).timezone() */

    	suite.T().Log("About to run line #117: r.Time(2013, r.July, 29).Timezone()")

        runAndAssert(suite.Suite, expected_, r.Time(2013, r.July, 29).Timezone(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #117")
    }

    {
        // times/api.yaml line #119
        /* 1375242965 */
        var expected_ int = 1375242965
        /* r.iso8601("2013-07-30T20:56:05-07:00").to_epoch_time() */

    	suite.T().Log("About to run line #119: r.ISO8601('2013-07-30T20:56:05-07:00').ToEpochTime()")

        runAndAssert(suite.Suite, expected_, r.ISO8601("2013-07-30T20:56:05-07:00").ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #119")
    }

    {
        // times/api.yaml line #122
        /* ("2013-07-30T20:56:05-07:00") */
        var expected_ string = "2013-07-30T20:56:05-07:00"
        /* r.epoch_time(1375242965).in_timezone("-07:00").to_iso8601() */

    	suite.T().Log("About to run line #122: r.EpochTime(1375242965).InTimezone('-07:00').ToISO8601()")

        runAndAssert(suite.Suite, expected_, r.EpochTime(1375242965).InTimezone("-07:00").ToISO8601(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #122")
    }

    {
        // times/api.yaml line #125
        /* ("PTYPE<TIME>") */
        var expected_ string = "PTYPE<TIME>"
        /* r.now().type_of() */

    	suite.T().Log("About to run line #125: r.Now().TypeOf()")

        runAndAssert(suite.Suite, expected_, r.Now().TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #125")
    }

    {
        // times/api.yaml line #127
        /* 0 */
        var expected_ int = 0
        /* (r.now() - r.now()) */

    	suite.T().Log("About to run line #127: r.Now().Sub(r.Now())")

        runAndAssert(suite.Suite, expected_, r.Now().Sub(r.Now()), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #127")
    }

    {
        // times/api.yaml line #132
        /* err("ReqlQueryLogicError", "ISO 8601 string has no time zone, and no default time zone was provided.", []) */
        var expected_ Err = err("ReqlQueryLogicError", "ISO 8601 string has no time zone, and no default time zone was provided.")
        /* r.iso8601("2013-07-30T20:56:05").to_iso8601() */

    	suite.T().Log("About to run line #132: r.ISO8601('2013-07-30T20:56:05').ToISO8601()")

        runAndAssert(suite.Suite, expected_, r.ISO8601("2013-07-30T20:56:05").ToISO8601(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #132")
    }

    {
        // times/api.yaml line #136
        /* ("2013-07-30T20:56:05-07:00") */
        var expected_ string = "2013-07-30T20:56:05-07:00"
        /* r.iso8601("2013-07-30T20:56:05", default_timezone='-07').to_iso8601() */

    	suite.T().Log("About to run line #136: r.ISO8601('2013-07-30T20:56:05', r.ISO8601Opts{DefaultTimezone: '-07', }).ToISO8601()")

        runAndAssert(suite.Suite, expected_, r.ISO8601("2013-07-30T20:56:05", r.ISO8601Opts{DefaultTimezone: "-07", }).ToISO8601(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #136")
    }

    {
        // times/api.yaml line #140
        /* ([1, 2, 3, 4, 5, 6, 7]) */
        var expected_ []interface{} = []interface{}{1, 2, 3, 4, 5, 6, 7}
        /* r.expr([r.monday, r.tuesday, r.wednesday, r.thursday, r.friday, r.saturday, r.sunday]) */

    	suite.T().Log("About to run line #140: r.Expr([]interface{}{r.Monday, r.Tuesday, r.Wednesday, r.Thursday, r.Friday, r.Saturday, r.Sunday})")

        runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{r.Monday, r.Tuesday, r.Wednesday, r.Thursday, r.Friday, r.Saturday, r.Sunday}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #140")
    }

    {
        // times/api.yaml line #142
        /* ([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]) */
        var expected_ []interface{} = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
        /* r.expr([r.january, r.february, r.march, r.april, r.may, r.june,
r.july, r.august, r.september, r.october, r.november, r.december]) */

    	suite.T().Log("About to run line #142: r.Expr([]interface{}{r.January, r.February, r.March, r.April, r.May, r.June, r.July, r.August, r.September, r.October, r.November, r.December})")

        runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{r.January, r.February, r.March, r.April, r.May, r.June, r.July, r.August, r.September, r.October, r.November, r.December}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
    	})
        suite.T().Log("Finished running line #142")
    }
}
