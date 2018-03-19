package main
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {

	db, err := sql.Open("mysql", "root@/teamtune")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO members(username, email, birthdate, password) VALUES(? , ? , ? , ?)")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec("threat", "vilius@rinkodara.lt", "1997-01-22", "5f4dcc3b5aa765d61d8327deb882cf99")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
