package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Получение статуса одной транзакции
func OneTransaction(id int) []TransactionStatus {
	db, err := sql.Open("mysql", dbparams) //открытие соединения с БД
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []TransactionStatus{}

	res, err := db.Query("SELECT `status` FROM `transactions` WHERE `id` = ?", id) // запрос к БД
	if err != nil {
		log.Fatal("No such transaction")
	}
	for res.Next() { // формирование сообщения результата
		var tr TransactionStatus
		err = res.Scan(&tr.Status)
		result = append(result, tr)
	}
	defer res.Close()
	return result
}

// Получение всех транзакций по ID пользователя
func AllTrasactionsId(userid int) Transactions {
	db, err := sql.Open("mysql", dbparams) //открытие соединения с БД
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []Transaction{}

	res, err := db.Query("SELECT * FROM `transactions` WHERE `userid` = ?", userid) // запрос к БД
	if err != nil {
		log.Fatal("No such user or transactions")
	}

	for res.Next() { // формирование сообщения результата
		var tr Transaction
		err = res.Scan(&tr.Id, &tr.Userid, &tr.Email, &tr.Price, &tr.Currency, &tr.CreatedOn, &tr.UpdatedOn, &tr.Status)
		result = append(result, tr)
	}
	defer res.Close()

	return result
}

// Получение всех транзакций по Email
func AllTransactionsEm(email string) Transactions {
	db, err := sql.Open("mysql", dbparams) //открытие соединения с БД
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []Transaction{}

	res, err := db.Query("SELECT * FROM `transactions` WHERE `email` = ?", email) // запрос к БД
	if err != nil {
		log.Fatal("No such user or transactions")
	}

	for res.Next() { // формирование сообщения результата
		var tr Transaction
		err = res.Scan(&tr.Id, &tr.Userid, &tr.Email, &tr.Price, &tr.Currency, &tr.CreatedOn, &tr.UpdatedOn, &tr.Status)
		result = append(result, tr)
	}
	defer res.Close()

	return result
}
