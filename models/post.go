package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddTransaction(t *Transaction) {
	db, err := sql.Open("mysql", dbparams)
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO `transactions` (`userid`, `email`, `price`, `currency`) VALUES (?, ?, ?, ?)", t.Userid, t.Email, t.Price, t.Currency)
	if err != nil {
		log.Fatal("Unable to add user")
	}
	defer insert.Close()
	log.Println("Транзакция успешно добавлена")
}
