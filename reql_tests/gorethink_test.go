package reql_tests

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	r "gopkg.in/dancannon/gorethink.v2"
)

var url string

func init() {
	flag.Parse()
	r.SetVerbose(true)

	// If the test is being run by wercker look for the rethink url
	url = os.Getenv("RETHINKDB_URL")
	if url == "" {
		url = "localhost:28015"
	}
}

// Print variable as JSON
func jsonPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}
