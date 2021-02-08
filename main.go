package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Gender    string
	Age       int32
}

type sqlUser struct {
	Conn *sql.DB
}

func NewSqlUser(Conn *sql.DB) *sqlUser {
	return &sqlUser{Conn}
}

func (s *sqlUser) Fetch(user *User) User {
	return User{
		ID: 0,
	}
}

func (s *sqlUser) Insert(user *User) {
	return
}

func (s *sqlUser) Update(user *User) {
	return
}

func (s *sqlUser) Delete(id int64) {
	return
}

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
