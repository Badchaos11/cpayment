package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello world!")
	Db, err := sql.Open("postgres", "host=localhost dbname=testdb user=user password=secret sslmode=disable")

	if err != nil {
		log.Fatal("Cannot connect to databaase")
	}
	defer Db.Close()
}
