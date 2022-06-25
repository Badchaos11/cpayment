package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Добавление транзакции в БД
func AddTransaction(t *Transaction) bool {
	db, err := sql.Open("mysql", dbparams) // подключени к БД
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()
	// Запрос к БД
	insert, err := db.Query("INSERT INTO `transactions` (`userid`, `email`, `price`, `currency`) VALUES (?, ?, ?, ?)", t.Userid, t.Email, t.Price, t.Currency)
	if err != nil {
		log.Fatal("Unable to add user")
	}
	defer insert.Close()
	log.Println("Транзакция успешно добавлена")
	// Возврат true для вывода сообщения об успехе
	return true
}

// Только как пример, добавление транзакции с ошибкой
func AddTransactionFail(t *Transaction) bool {
	db, err := sql.Open("mysql", dbparams) // подключени к БД
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()
	// Запрос к БД
	insert, err := db.Query("INSERT INTO `transactions` (`userid`, `email`, `price`, `currency`, `status`) VALUES (?, ?, ?, ?, ?)", t.Userid, t.Email, t.Price, t.Currency, t.Status)
	if err != nil {
		log.Fatal("Unable to add user")
	}
	defer insert.Close()
	log.Println("При добавлении транзакции произошла ошибка")
	// Возврат true для вывода сообщения об успехе
	return true
}
