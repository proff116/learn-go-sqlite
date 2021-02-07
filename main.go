package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*type User struct {
	FirstName string
	LastName  string
	Gender    string
	Age       int32
}*/

func main() {
	db, err := sql.Open("sqlite3", "users_db.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	result, err := db.Exec("insert into Users (first_name, last_name, gender, age) values ('firstName', 'lastName', 'Male', 18)")

	if result == nil {
		panic(result)
	}

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select * from Users")

	if err != nil {
		panic(err)
	}

	defer rows.Close()
}
