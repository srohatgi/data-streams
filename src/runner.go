package main

import (
	"github.com/araddon/cass"
	"fmt"
	"net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}

func main() {
	fmt.Printf("hello world 2")

	cass.ConfigKeyspace("testing", []string{"127.0.0.1:9160"}, 20)
	conn, _ := cass.GetCassConn("testing")

	defer conn.Checkin() // there is a pool of connections, so you must checkin/out

	// 1 is replication factor, will return err if already exist
	_ = conn.CreateKeyspace("testing", 1)

	_ = conn.CreateCF("col_fam_name")

	var cols = map[string]string{
		"lastname": "golang",
	}

	// if you pass 0 timestamp it will generate one
	err := conn.Insert("col_fam_name", "keyinserttest", cols, 0)
	if err != nil {
		fmt.Println("error, insert/read failed")
	}

	
	col, _ := conn.Get("col_fam_name", "keyinserttest", "lastname")
	if col == nil {
		fmt.Println("insert/get single row, single col failed: testcol - keyinserttest")
	}

	//http.HandleFunc("/", defaultHandler)
	//http.ListenAndServe(":8080", nil)

}
