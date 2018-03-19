package main

import (
	"strconv"
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Payload struct {
	ErrorMessage string
	ReturnCode int
}

func main() {
	http.HandleFunc("/create", Create)
	http.ListenAndServe("localhost:1339", nil)
}



func checkUserExists() {

}

func Create(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root@/teamtune")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	r.ParseForm()                     // Parses the request body

	name := r.Form.Get("name") // x will be "" if parameter is not set
	description := r.Form.Get("description") // x will be "" if parameter is not set
	member := 1

	ErrorMsg := ""
	ErrorCode := 0

	p := Payload{ErrorMsg, ErrorCode}

	// Execute the query
	rows, err := db.Query("SELECT id FROM members WHERE id = " + strconv.Itoa(member))
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
	}


	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO projects(admin, name, description) VALUES(? , ? , ?)")
	if err != nil {
		ErrorMsg = "Error generating query"
		ErrorCode = 301
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	result, err := stmtIns.Exec(member, name, description)
	if err != nil {
		ErrorMsg = "Error inserting project row"
		ErrorCode = 300
	} else
	{
		id, err := result.LastInsertId()
		if err != nil {
			ErrorMsg = "Error parsing insertionID"
			ErrorCode = 300
		}
		ErrorMsg = "Project Created. ID - " + strconv.FormatInt(id, 32)
	}
	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil{
		panic(nil)
	}
	fmt.Fprintf(w, string(json))
}