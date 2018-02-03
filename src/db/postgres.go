package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"time"
)

var Db *sql.DB

func main() {
	newUser:=User{
		Id:1,
		Name:"name",
		Uuid:"001",
		Email:"sunth@126.com",
		Password:"12",
	}

	newUser.Create()
}

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat user=postgres password=sunfei sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (user *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow("001",user.Name, user.Email, user.Password, time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}
