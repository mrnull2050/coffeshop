package database

import (
	
	"database/sql"
	"fmt"
	"log"
)


func Connected() *sql.DB{
	//connect to db
	dsn := "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable"
	db , err := sql.Open("postgres" , dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected database")
	return db
}