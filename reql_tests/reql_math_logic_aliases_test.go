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

// Test named aliases for math and logic operators
func TestMathLogicAliasesSuite(t *testing.T) {
	suite.Run(t, new(MathLogicAliasesSuite ))
}

type MathLogicAliasesSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MathLogicAliasesSuite) SetupTest() {
	suite.T().Log("Setting up MathLogicAliasesSuite")
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

func (suite *MathLogicAliasesSuite) TearDownSuite() {
	suite.T().Log("Tearing down MathLogicAliasesSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MathLogicAliasesSuite) TestCases() {
	suite.T().Log("Running MathLogicAliasesSuite: Test named aliases for math and logic operators")



	{
		// math_logic/aliases.yaml line #5
		/* 1 */
		var expected_ int = 1
		/* r.expr(0).add(1) */

		suite.T().Log("About to run line #5: r.Expr(0).Add(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(0).Add(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #5")
	}

	{
		// math_logic/aliases.yaml line #6
		/* 1 */
		var expected_ int = 1
		/* r.add(0, 1) */

		suite.T().Log("About to run line #6: r.Add(0, 1)")

		runAndAssert(suite.Suite, expected_, r.Add(0, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// math_logic/aliases.yaml line #7
		/* 1 */
		var expected_ int = 1
		/* r.expr(2).sub(1) */

		suite.T().Log("About to run line #7: r.Expr(2).Sub(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(2).Sub(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// math_logic/aliases.yaml line #8
		/* 1 */
		var expected_ int = 1
		/* r.sub(2, 1) */

		suite.T().Log("About to run line #8: r.Sub(2, 1)")

		runAndAssert(suite.Suite, expected_, r.Sub(2, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// math_logic/aliases.yaml line #9
		/* 1 */
		var expected_ int = 1
		/* r.expr(2).div(2) */

		suite.T().Log("About to run line #9: r.Expr(2).Div(2)")

		runAndAssert(suite.Suite, expected_, r.Expr(2).Div(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #9")
	}

	{
		// math_logic/aliases.yaml line #10
		/* 1 */
		var expected_ int = 1
		/* r.div(2, 2) */

		suite.T().Log("About to run line #10: r.Div(2, 2)")

		runAndAssert(suite.Suite, expected_, r.Div(2, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// math_logic/aliases.yaml line #11
		/* 1 */
		var expected_ int = 1
		/* r.expr(1).mul(1) */

		suite.T().Log("About to run line #11: r.Expr(1).Mul(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Mul(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #11")
	}

	{
		// math_logic/aliases.yaml line #12
		/* 1 */
		var expected_ int = 1
		/* r.mul(1, 1) */

		suite.T().Log("About to run line #12: r.Mul(1, 1)")

		runAndAssert(suite.Suite, expected_, r.Mul(1, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #12")
	}

	{
		// math_logic/aliases.yaml line #13
		/* 1 */
		var expected_ int = 1
		/* r.expr(1).mod(2) */

		suite.T().Log("About to run line #13: r.Expr(1).Mod(2)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Mod(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #13")
	}

	{
		// math_logic/aliases.yaml line #14
		/* 1 */
		var expected_ int = 1
		/* r.mod(1, 2) */

		suite.T().Log("About to run line #14: r.Mod(1, 2)")

		runAndAssert(suite.Suite, expected_, r.Mod(1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #14")
	}

	{
		// math_logic/aliases.yaml line #25
		/* True */
		var expected_ bool = true
		/* r.expr(True).and_(True) */

		suite.T().Log("About to run line #25: r.Expr(true).And(true)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).And(true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #25")
	}

	{
		// math_logic/aliases.yaml line #26
		/* True */
		var expected_ bool = true
		/* r.expr(True).or_(True) */

		suite.T().Log("About to run line #26: r.Expr(true).Or(true)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).Or(true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #26")
	}

	{
		// math_logic/aliases.yaml line #27
		/* True */
		var expected_ bool = true
		/* r.and_(True, True) */

		suite.T().Log("About to run line #27: r.And(true, true)")

		runAndAssert(suite.Suite, expected_, r.And(true, true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #27")
	}

	{
		// math_logic/aliases.yaml line #28
		/* True */
		var expected_ bool = true
		/* r.or_(True, True) */

		suite.T().Log("About to run line #28: r.Or(true, true)")

		runAndAssert(suite.Suite, expected_, r.Or(true, true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// math_logic/aliases.yaml line #29
		/* True */
		var expected_ bool = true
		/* r.expr(False).not_() */

		suite.T().Log("About to run line #29: r.Expr(false).Not()")

		runAndAssert(suite.Suite, expected_, r.Expr(false).Not(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #29")
	}

	{
		// math_logic/aliases.yaml line #30
		/* True */
		var expected_ bool = true
		/* r.not_(False) */

		suite.T().Log("About to run line #30: r.Not(false)")

		runAndAssert(suite.Suite, expected_, r.Not(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// math_logic/aliases.yaml line #34
		/* True */
		var expected_ bool = true
		/* r.expr(1).eq(1) */

		suite.T().Log("About to run line #34: r.Expr(1).Eq(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Eq(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #34")
	}

	{
		// math_logic/aliases.yaml line #35
		/* True */
		var expected_ bool = true
		/* r.expr(1).ne(2) */

		suite.T().Log("About to run line #35: r.Expr(1).Ne(2)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Ne(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #35")
	}

	{
		// math_logic/aliases.yaml line #36
		/* True */
		var expected_ bool = true
		/* r.expr(1).lt(2) */

		suite.T().Log("About to run line #36: r.Expr(1).Lt(2)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Lt(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #36")
	}

	{
		// math_logic/aliases.yaml line #37
		/* True */
		var expected_ bool = true
		/* r.expr(1).gt(0) */

		suite.T().Log("About to run line #37: r.Expr(1).Gt(0)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Gt(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// math_logic/aliases.yaml line #38
		/* True */
		var expected_ bool = true
		/* r.expr(1).le(1) */

		suite.T().Log("About to run line #38: r.Expr(1).Le(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Le(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #38")
	}

	{
		// math_logic/aliases.yaml line #39
		/* True */
		var expected_ bool = true
		/* r.expr(1).ge(1) */

		suite.T().Log("About to run line #39: r.Expr(1).Ge(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Ge(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #39")
	}

	{
		// math_logic/aliases.yaml line #40
		/* True */
		var expected_ bool = true
		/* r.eq(1, 1) */

		suite.T().Log("About to run line #40: r.Eq(1, 1)")

		runAndAssert(suite.Suite, expected_, r.Eq(1, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #40")
	}

	{
		// math_logic/aliases.yaml line #41
		/* True */
		var expected_ bool = true
		/* r.ne(1, 2) */

		suite.T().Log("About to run line #41: r.Ne(1, 2)")

		runAndAssert(suite.Suite, expected_, r.Ne(1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #41")
	}

	{
		// math_logic/aliases.yaml line #42
		/* True */
		var expected_ bool = true
		/* r.lt(1, 2) */

		suite.T().Log("About to run line #42: r.Lt(1, 2)")

		runAndAssert(suite.Suite, expected_, r.Lt(1, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #42")
	}

	{
		// math_logic/aliases.yaml line #43
		/* True */
		var expected_ bool = true
		/* r.gt(1, 0) */

		suite.T().Log("About to run line #43: r.Gt(1, 0)")

		runAndAssert(suite.Suite, expected_, r.Gt(1, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #43")
	}

	{
		// math_logic/aliases.yaml line #44
		/* True */
		var expected_ bool = true
		/* r.le(1, 1) */

		suite.T().Log("About to run line #44: r.Le(1, 1)")

		runAndAssert(suite.Suite, expected_, r.Le(1, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #44")
	}

	{
		// math_logic/aliases.yaml line #45
		/* True */
		var expected_ bool = true
		/* r.ge(1, 1) */

		suite.T().Log("About to run line #45: r.Ge(1, 1)")

		runAndAssert(suite.Suite, expected_, r.Ge(1, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
		})
		suite.T().Log("Finished running line #45")
	}
}
