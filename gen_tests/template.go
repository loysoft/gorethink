package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/dancannon/gorethink.v2"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func Test${module_name}Suite(t *testing.T) {
    suite.Run(t, new(${module_name}Suite ))
}

type ${module_name}Suite struct {
	suite.Suite

	session *r.Session
}

func (suite *${module_name}Suite) SetupTest() {
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

	%for var_name in table_var_names:
	r.DB("test").TableDrop("${var_name}").Exec(suite.session)
	err = r.DB("test").TableCreate("${var_name}").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("${var_name}").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	%endfor
}

func (suite *${module_name}Suite) TearDownSuite() {
	r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
    %for var_name in table_var_names:
	 r.DB("test").TableDrop("${var_name}").Exec(suite.session)
    %endfor
    r.DBDrop("test").Exec(suite.session)

    suite.session.Close()
}

<%rendered_vars = set() %>\
func (suite *${module_name}Suite) TestCases() {
	%for var_name in table_var_names:
	${var_name} := r.DB("test").Table("${var_name}")
	%endfor

<%rendered_something = False %>\
    %for item in defs_and_test:
    %if type(item) == GoDef:
<%rendered_something = True %>
    // ${item.testfile} line #${item.line_num}
    // ${item.line.original.replace('\n', '')}
    suite.T().Log("Possibly executing: ${item.line.go.replace('\\', '\\\\').replace('"', "'")}")

    %if item.varname in rendered_vars:
    ${item.varname} = ${item.value}
    %else:
    var ${item.varname} ${item.vartype} = ${item.value}
	<%rendered_vars.add(item.varname)%>\
    %endif

    %elif type(item) == GoQuery:
<%rendered_something = True %>
    {
        // ${item.testfile} line #${item.line_num}
        /* ${item.expected_line.original} */
        var expected_ ${item.expected_type} = ${item.expected_line.go}
        /* ${item.line.original} */

    	suite.T().Log("About to run line #${item.line_num}: ${item.line.go.replace('"', "'").replace('\\', '\\\\').replace('\n', '\\n')}")

    	%if item.line.go.startswith('fetch(') and item.line.go.endswith(')'):
        fetch(suite.Suite, suite.session, expected_, ${item.line.go[6:-1]}, r.RunOpts{
			%if item.runopts:
			%for key, val in item.runopts.items():
			${key}: ${val},
			%endfor
			%endif
    	})
	    %elif item.is_value:
        actual := ${item.line.go}

    	assertCompare(suite.T(), expected_, actual)
    	%else:
        cursor, err := ${item.line.go}.Run(suite.session, r.RunOpts{
			%if item.runopts:
			%for key, val in item.runopts.items():
			${key}: ${val},
			%endfor
			%endif
    	})

    	assertExpected(suite.Suite, expected_, cursor, err)
    	%endif
        suite.T().Log("Finished running line #${item.line_num}")
    }
    %endif
    %endfor
    %if not rendered_something:
<% raise EmptyTemplate() %>\
    %endif
}
