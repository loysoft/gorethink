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

// Test edge cases of changefeed operations
func TestChangefeedsEdgeSuite(t *testing.T) {
	suite.Run(t, new(ChangefeedsEdgeSuite ))
}

type ChangefeedsEdgeSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *ChangefeedsEdgeSuite) SetupTest() {
	suite.T().Log("Setting up ChangefeedsEdgeSuite")
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

	r.DB("test").TableDrop("tbl").Exec(suite.session)
	err = r.DB("test").TableCreate("tbl").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("tbl").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *ChangefeedsEdgeSuite) TearDownSuite() {
	suite.T().Log("Tearing down ChangefeedsEdgeSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		 r.DB("test").TableDrop("tbl").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *ChangefeedsEdgeSuite) TestCases() {
	suite.T().Log("Running ChangefeedsEdgeSuite: Test edge cases of changefeed operations")

	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors


	// changefeeds/edge.yaml line #5
	// common_prefix = r.expr([0,1,2,3,4,5,6,7,8])
	suite.T().Log("Possibly executing: var common_prefix r.Term = r.Expr([]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8})")

	common_prefix := r.Expr([]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8})
	_ = common_prefix // Prevent any noused variable errors


	{
		// changefeeds/edge.yaml line #8
		/* ({'created':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1, }
		/* tbl.index_create('sindex', lambda row:common_prefix.append(row['value'])) */

		suite.T().Log("About to run line #8: tbl.IndexCreateFunc('sindex', func(row r.Term) interface{} { return common_prefix.Append(row.AtIndex('value'))})")

		runAndAssert(suite.Suite, expected_, tbl.IndexCreateFunc("sindex", func(row r.Term) interface{} { return common_prefix.Append(row.AtIndex("value"))}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// changefeeds/edge.yaml line #11
		/* AnythingIsFine */
		var expected_ string = AnythingIsFine
		/* tbl.index_wait('sindex') */

		suite.T().Log("About to run line #11: tbl.IndexWait('sindex')")

		runAndAssert(suite.Suite, expected_, tbl.IndexWait("sindex"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #11")
	}

	// changefeeds/edge.yaml line #14
	// pre = r.range(7).coerce_to('array').add(r.range(10,70).coerce_to('array')).append(100).map(r.row.coerce_to('string'))
	suite.T().Log("Possibly executing: var pre r.Term = r.Range(7).CoerceTo('array').Add(r.Range(10, 70).CoerceTo('array')).Append(100).Map(r.Row.CoerceTo('string'))")

	pre := maybeRun(r.Range(7).CoerceTo("array").Add(r.Range(10, 70).CoerceTo("array")).Append(100).Map(r.Row.CoerceTo("string")), suite.session, r.RunOpts{
	});
	_ = pre // Prevent any noused variable errors


	// changefeeds/edge.yaml line #16
	// mid = r.range(2,9).coerce_to('array').add(r.range(20,90).coerce_to('array')).map(r.row.coerce_to('string'))
	suite.T().Log("Possibly executing: var mid r.Term = r.Range(2, 9).CoerceTo('array').Add(r.Range(20, 90).CoerceTo('array')).Map(r.Row.CoerceTo('string'))")

	mid := maybeRun(r.Range(2, 9).CoerceTo("array").Add(r.Range(20, 90).CoerceTo("array")).Map(r.Row.CoerceTo("string")), suite.session, r.RunOpts{
	});
	_ = mid // Prevent any noused variable errors


	// changefeeds/edge.yaml line #18
	// post = r.range(3,10).coerce_to('array').add(r.range(30,100).coerce_to('array')).map(r.row.coerce_to('string'))
	suite.T().Log("Possibly executing: var post r.Term = r.Range(3, 10).CoerceTo('array').Add(r.Range(30, 100).CoerceTo('array')).Map(r.Row.CoerceTo('string'))")

	post := maybeRun(r.Range(3, 10).CoerceTo("array").Add(r.Range(30, 100).CoerceTo("array")).Map(r.Row.CoerceTo("string")), suite.session, r.RunOpts{
	});
	_ = post // Prevent any noused variable errors


	// changefeeds/edge.yaml line #21
	// erroredres = r.range(2).coerce_to('array').add(r.range(10, 20).coerce_to('array')).append(100).map(r.row.coerce_to('string'))
	suite.T().Log("Possibly executing: var erroredres r.Term = r.Range(2).CoerceTo('array').Add(r.Range(10, 20).CoerceTo('array')).Append(100).Map(r.Row.CoerceTo('string'))")

	erroredres := maybeRun(r.Range(2).CoerceTo("array").Add(r.Range(10, 20).CoerceTo("array")).Append(100).Map(r.Row.CoerceTo("string")), suite.session, r.RunOpts{
	});
	_ = erroredres // Prevent any noused variable errors


	// changefeeds/edge.yaml line #26
	// pre_changes = tbl.between(r.minval, common_prefix.append('7'), index='sindex').changes(squash=False).limit(len(pre))['new_val']['value']
	suite.T().Log("Possibly executing: var pre_changes r.Term = tbl.Between(r.MinVal, common_prefix.Append('7'), r.BetweenOpts{Index: 'sindex', }).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(pre)).AtIndex('new_val').AtIndex('value')")

	pre_changes := maybeRun(tbl.Between(r.MinVal, common_prefix.Append("7"), r.BetweenOpts{Index: "sindex", }).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(pre)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = pre_changes // Prevent any noused variable errors


	// changefeeds/edge.yaml line #29
	// mid_changes = tbl.between(common_prefix.append('2'), common_prefix.append('9'), index='sindex').changes(squash=False).limit(len(post))['new_val']['value']
	suite.T().Log("Possibly executing: var mid_changes r.Term = tbl.Between(common_prefix.Append('2'), common_prefix.Append('9'), r.BetweenOpts{Index: 'sindex', }).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(post)).AtIndex('new_val').AtIndex('value')")

	mid_changes := maybeRun(tbl.Between(common_prefix.Append("2"), common_prefix.Append("9"), r.BetweenOpts{Index: "sindex", }).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(post)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = mid_changes // Prevent any noused variable errors


	// changefeeds/edge.yaml line #32
	// post_changes = tbl.between(common_prefix.append('3'), r.maxval, index='sindex').changes(squash=False).limit(len(mid))['new_val']['value']
	suite.T().Log("Possibly executing: var post_changes r.Term = tbl.Between(common_prefix.Append('3'), r.MaxVal, r.BetweenOpts{Index: 'sindex', }).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(mid)).AtIndex('new_val').AtIndex('value')")

	post_changes := maybeRun(tbl.Between(common_prefix.Append("3"), r.MaxVal, r.BetweenOpts{Index: "sindex", }).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(mid)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = post_changes // Prevent any noused variable errors


	// changefeeds/edge.yaml line #38
	// premap_changes1 = tbl.map(r.branch(r.row['value'].lt('2'), r.row, r.row["dummy"])).changes(squash=False).limit(len(erroredres))['new_val']['value']
	suite.T().Log("Possibly executing: var premap_changes1 r.Term = tbl.Map(r.Branch(r.Row.AtIndex('value').Lt('2'), r.Row, r.Row.AtIndex('dummy'))).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(erroredres)).AtIndex('new_val').AtIndex('value')")

	premap_changes1 := maybeRun(tbl.Map(r.Branch(r.Row.AtIndex("value").Lt("2"), r.Row, r.Row.AtIndex("dummy"))).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(erroredres)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = premap_changes1 // Prevent any noused variable errors


	// changefeeds/edge.yaml line #42
	// postmap_changes1 = tbl.changes(squash=False).map(r.branch(r.row['new_val']['value'].lt('2'), r.row, r.row["dummy"])).limit(len(erroredres))['new_val']['value']
	suite.T().Log("Possibly executing: var postmap_changes1 r.Term = tbl.Changes(r.ChangesOpts{Squash: false, }).Map(r.Branch(r.Row.AtIndex('new_val').AtIndex('value').Lt('2'), r.Row, r.Row.AtIndex('dummy'))).Limit(maybeLen(erroredres)).AtIndex('new_val').AtIndex('value')")

	postmap_changes1 := maybeRun(tbl.Changes(r.ChangesOpts{Squash: false, }).Map(r.Branch(r.Row.AtIndex("new_val").AtIndex("value").Lt("2"), r.Row, r.Row.AtIndex("dummy"))).Limit(maybeLen(erroredres)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = postmap_changes1 // Prevent any noused variable errors


	// changefeeds/edge.yaml line #46
	// prefilter_changes1 = tbl.filter(r.branch(r.row['value'].lt('2'), True, r.row["dummy"])).changes(squash=False).limit(len(erroredres))['new_val']['value']
	suite.T().Log("Possibly executing: var prefilter_changes1 r.Term = tbl.Filter(r.Branch(r.Row.AtIndex('value').Lt('2'), true, r.Row.AtIndex('dummy'))).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(erroredres)).AtIndex('new_val').AtIndex('value')")

	prefilter_changes1 := maybeRun(tbl.Filter(r.Branch(r.Row.AtIndex("value").Lt("2"), true, r.Row.AtIndex("dummy"))).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(erroredres)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = prefilter_changes1 // Prevent any noused variable errors


	// changefeeds/edge.yaml line #50
	// postfilter_changes1 = tbl.changes(squash=False).filter(r.branch(r.row['new_val']['value'].lt('2'), True, r.row["dummy"])).limit(len(erroredres))['new_val']['value']
	suite.T().Log("Possibly executing: var postfilter_changes1 r.Term = tbl.Changes(r.ChangesOpts{Squash: false, }).Filter(r.Branch(r.Row.AtIndex('new_val').AtIndex('value').Lt('2'), true, r.Row.AtIndex('dummy'))).Limit(maybeLen(erroredres)).AtIndex('new_val').AtIndex('value')")

	postfilter_changes1 := maybeRun(tbl.Changes(r.ChangesOpts{Squash: false, }).Filter(r.Branch(r.Row.AtIndex("new_val").AtIndex("value").Lt("2"), true, r.Row.AtIndex("dummy"))).Limit(maybeLen(erroredres)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = postfilter_changes1 // Prevent any noused variable errors


	// changefeeds/edge.yaml line #56
	// premap_changes2 = tbl.map(r.branch(r.row['value'].lt('2'), r.row, r.expr([])[1])).changes(squash=False).limit(len(erroredres))['new_val']['value']
	suite.T().Log("Possibly executing: var premap_changes2 r.Term = tbl.Map(r.Branch(r.Row.AtIndex('value').Lt('2'), r.Row, r.Expr([]interface{}{}).AtIndex(1))).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(erroredres)).AtIndex('new_val').AtIndex('value')")

	premap_changes2 := maybeRun(tbl.Map(r.Branch(r.Row.AtIndex("value").Lt("2"), r.Row, r.Expr([]interface{}{}).AtIndex(1))).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(erroredres)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = premap_changes2 // Prevent any noused variable errors


	// changefeeds/edge.yaml line #60
	// postmap_changes2 = tbl.changes(squash=False).map(r.branch(r.row['new_val']['value'].lt('2'), r.row, r.expr([])[1])).limit(len(erroredres))['new_val']['value']
	suite.T().Log("Possibly executing: var postmap_changes2 r.Term = tbl.Changes(r.ChangesOpts{Squash: false, }).Map(r.Branch(r.Row.AtIndex('new_val').AtIndex('value').Lt('2'), r.Row, r.Expr([]interface{}{}).AtIndex(1))).Limit(maybeLen(erroredres)).AtIndex('new_val').AtIndex('value')")

	postmap_changes2 := maybeRun(tbl.Changes(r.ChangesOpts{Squash: false, }).Map(r.Branch(r.Row.AtIndex("new_val").AtIndex("value").Lt("2"), r.Row, r.Expr([]interface{}{}).AtIndex(1))).Limit(maybeLen(erroredres)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = postmap_changes2 // Prevent any noused variable errors


	// changefeeds/edge.yaml line #64
	// prefilter_changes2 = tbl.filter(r.branch(r.row['value'].lt('2'), True, r.expr([])[1])).changes(squash=False).limit(len(erroredres))['new_val']['value']
	suite.T().Log("Possibly executing: var prefilter_changes2 r.Term = tbl.Filter(r.Branch(r.Row.AtIndex('value').Lt('2'), true, r.Expr([]interface{}{}).AtIndex(1))).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(erroredres)).AtIndex('new_val').AtIndex('value')")

	prefilter_changes2 := maybeRun(tbl.Filter(r.Branch(r.Row.AtIndex("value").Lt("2"), true, r.Expr([]interface{}{}).AtIndex(1))).Changes(r.ChangesOpts{Squash: false, }).Limit(maybeLen(erroredres)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = prefilter_changes2 // Prevent any noused variable errors


	// changefeeds/edge.yaml line #68
	// postfilter_changes2 = tbl.changes(squash=False).filter(r.branch(r.row['new_val']['value'].lt('2'), True, r.expr([])[1])).limit(len(erroredres))['new_val']['value']
	suite.T().Log("Possibly executing: var postfilter_changes2 r.Term = tbl.Changes(r.ChangesOpts{Squash: false, }).Filter(r.Branch(r.Row.AtIndex('new_val').AtIndex('value').Lt('2'), true, r.Expr([]interface{}{}).AtIndex(1))).Limit(maybeLen(erroredres)).AtIndex('new_val').AtIndex('value')")

	postfilter_changes2 := maybeRun(tbl.Changes(r.ChangesOpts{Squash: false, }).Filter(r.Branch(r.Row.AtIndex("new_val").AtIndex("value").Lt("2"), true, r.Expr([]interface{}{}).AtIndex(1))).Limit(maybeLen(erroredres)).AtIndex("new_val").AtIndex("value"), suite.session, r.RunOpts{
	});
	_ = postfilter_changes2 // Prevent any noused variable errors


	// changefeeds/edge.yaml line #73
	// nondetermmap = r.branch(r.random().gt(0.5), r.row, r.error("dummy"))
	suite.T().Log("Possibly executing: var nondetermmap r.Term = r.Branch(r.Random().Gt(0.5), r.Row, r.Error('dummy'))")

	nondetermmap := r.Branch(r.Random().Gt(0.5), r.Row, r.Error("dummy"))
	_ = nondetermmap // Prevent any noused variable errors


	// changefeeds/edge.yaml line #77
	// nondetermfilter = lambda :r.random().gt(0.5)
	suite.T().Log("Possibly executing: var nondetermfilter func() = func() interface{} { return r.Random().Gt(0.5)}")

	nondetermfilter := func() interface{} { return r.Random().Gt(0.5)}
	_ = nondetermfilter // Prevent any noused variable errors


	{
		// changefeeds/edge.yaml line #83
		/* err('ReqlQueryLogicError', 'Cannot call `changes` after a non-deterministic function.') */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot call `changes` after a non-deterministic function.")
		/* tbl.map(nondetermmap).changes(squash=False) */

		suite.T().Log("About to run line #83: tbl.Map(nondetermmap).Changes(r.ChangesOpts{Squash: false, })")

		runAndAssert(suite.Suite, expected_, tbl.Map(nondetermmap).Changes(r.ChangesOpts{Squash: false, }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #83")
	}

	// changefeeds/edge.yaml line #88
	// postmap_changes3 = tbl.changes(squash=False).map(nondetermmap).limit(100)
	suite.T().Log("Possibly executing: var postmap_changes3 r.Term = tbl.Changes(r.ChangesOpts{Squash: false, }).Map(nondetermmap).Limit(100)")

	postmap_changes3 := maybeRun(tbl.Changes(r.ChangesOpts{Squash: false, }).Map(nondetermmap).Limit(100), suite.session, r.RunOpts{
	});
	_ = postmap_changes3 // Prevent any noused variable errors


	{
		// changefeeds/edge.yaml line #92
		/* err('ReqlQueryLogicError', 'Cannot call `changes` after a non-deterministic function.') */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot call `changes` after a non-deterministic function.")
		/* tbl.filter(nondetermfilter).changes(squash=False) */

		suite.T().Log("About to run line #92: tbl.Filter(nondetermfilter).Changes(r.ChangesOpts{Squash: false, })")

		runAndAssert(suite.Suite, expected_, tbl.Filter(nondetermfilter).Changes(r.ChangesOpts{Squash: false, }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #92")
	}

	// changefeeds/edge.yaml line #97
	// postfilter_changes3 = tbl.changes(squash=False).filter(nondetermfilter).limit(4)
	suite.T().Log("Possibly executing: var postfilter_changes3 r.Term = tbl.Changes(r.ChangesOpts{Squash: false, }).Filter(nondetermfilter).Limit(4)")

	postfilter_changes3 := maybeRun(tbl.Changes(r.ChangesOpts{Squash: false, }).Filter(nondetermfilter).Limit(4), suite.session, r.RunOpts{
	});
	_ = postfilter_changes3 // Prevent any noused variable errors


	{
		// changefeeds/edge.yaml line #100
		/* ({'skipped':0,'deleted':0,'unchanged':0,'errors':0,'replaced':0,'inserted':101}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"skipped": 0, "deleted": 0, "unchanged": 0, "errors": 0, "replaced": 0, "inserted": 101, }
		/* tbl.insert(r.range(101).map({'id':r.uuid().coerce_to('binary').slice(0,r.random(4,24)).coerce_to('string'),'value':r.row.coerce_to('string')})) */

		suite.T().Log("About to run line #100: tbl.Insert(r.Range(101).Map(map[interface{}]interface{}{'id': r.UUID().CoerceTo('binary').Slice(0, r.Random(4, 24)).CoerceTo('string'), 'value': r.Row.CoerceTo('string'), }))")

		runAndAssert(suite.Suite, expected_, tbl.Insert(r.Range(101).Map(map[interface{}]interface{}{"id": r.UUID().CoerceTo("binary").Slice(0, r.Random(4, 24)).CoerceTo("string"), "value": r.Row.CoerceTo("string"), })), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #100")
	}

	{
		// changefeeds/edge.yaml line #105
		/* bag(pre) */
		var expected_ Expected = bag(pre)
		/* pre_changes */

		suite.T().Log("About to run line #105: pre_changes")

		runAndAssert(suite.Suite, expected_, pre_changes, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #105")
	}

	{
		// changefeeds/edge.yaml line #108
		/* bag(mid) */
		var expected_ Expected = bag(mid)
		/* mid_changes */

		suite.T().Log("About to run line #108: mid_changes")

		runAndAssert(suite.Suite, expected_, mid_changes, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #108")
	}

	{
		// changefeeds/edge.yaml line #111
		/* bag(post) */
		var expected_ Expected = bag(post)
		/* post_changes */

		suite.T().Log("About to run line #111: post_changes")

		runAndAssert(suite.Suite, expected_, post_changes, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #111")
	}

	{
		// changefeeds/edge.yaml line #114
		/* bag(erroredres) */
		var expected_ Expected = bag(erroredres)
		/* premap_changes1 */

		suite.T().Log("About to run line #114: premap_changes1")

		runAndAssert(suite.Suite, expected_, premap_changes1, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #114")
	}

	{
		// changefeeds/edge.yaml line #117
		/* bag(erroredres) */
		var expected_ Expected = bag(erroredres)
		/* premap_changes2 */

		suite.T().Log("About to run line #117: premap_changes2")

		runAndAssert(suite.Suite, expected_, premap_changes2, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #117")
	}

	{
		// changefeeds/edge.yaml line #120
		/* err('ReqlNonExistenceError', "No attribute `dummy` in object:") */
		var expected_ Err = err("ReqlNonExistenceError", "No attribute `dummy` in object:")
		/* postmap_changes1 */

		suite.T().Log("About to run line #120: postmap_changes1")

		runAndAssert(suite.Suite, expected_, postmap_changes1, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #120")
	}

	{
		// changefeeds/edge.yaml line #123
		/* err('ReqlNonExistenceError', "Index out of bounds:" + " 1") */
		var expected_ Err = err("ReqlNonExistenceError", "Index out of bounds:" + " 1")
		/* postmap_changes2 */

		suite.T().Log("About to run line #123: postmap_changes2")

		runAndAssert(suite.Suite, expected_, postmap_changes2, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #123")
	}

	{
		// changefeeds/edge.yaml line #126
		/* err('ReqlUserError', "dummy") */
		var expected_ Err = err("ReqlUserError", "dummy")
		/* postmap_changes3 */

		suite.T().Log("About to run line #126: postmap_changes3")

		runAndAssert(suite.Suite, expected_, postmap_changes3, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #126")
	}

	{
		// changefeeds/edge.yaml line #129
		/* bag(erroredres) */
		var expected_ Expected = bag(erroredres)
		/* prefilter_changes1 */

		suite.T().Log("About to run line #129: prefilter_changes1")

		runAndAssert(suite.Suite, expected_, prefilter_changes1, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #129")
	}

	{
		// changefeeds/edge.yaml line #132
		/* bag(erroredres) */
		var expected_ Expected = bag(erroredres)
		/* prefilter_changes2 */

		suite.T().Log("About to run line #132: prefilter_changes2")

		runAndAssert(suite.Suite, expected_, prefilter_changes2, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #132")
	}

	{
		// changefeeds/edge.yaml line #135
		/* bag(erroredres) */
		var expected_ Expected = bag(erroredres)
		/* postfilter_changes1 */

		suite.T().Log("About to run line #135: postfilter_changes1")

		runAndAssert(suite.Suite, expected_, postfilter_changes1, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #135")
	}

	{
		// changefeeds/edge.yaml line #138
		/* bag(erroredres) */
		var expected_ Expected = bag(erroredres)
		/* postfilter_changes2 */

		suite.T().Log("About to run line #138: postfilter_changes2")

		runAndAssert(suite.Suite, expected_, postfilter_changes2, suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #138")
	}
}
