package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddTransaction(t *Transaction) bool {
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
	return true
}

func AddTransactionFail(t *Transaction) bool {
	db, err := sql.Open("mysql", dbparams)
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO `transactions` (`userid`, `email`, `price`, `currency`, `status`) VALUES (?, ?, ?, ?, ?)", t.Userid, t.Email, t.Price, t.Currency, t.Status)
	if err != nil {
		log.Fatal("Unable to add user")
	}
	defer insert.Close()
	log.Println("При добавлении транзакции произошла ошибка")
	return true
}
