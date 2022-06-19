package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OneTransaction(id int) []TransactionStatus {
	db, err := sql.Open("mysql", dbparams)
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []TransactionStatus{}

	res, err := db.Query("SELECT `status` FROM `transactions` WHERE `id` = ?", id)
	if err != nil {
		log.Fatal("No such transaction")
	}
	for res.Next() {
		var tr TransactionStatus
		err = res.Scan(&tr.Status)
		result = append(result, tr)
	}
	defer res.Close()
	return result
}

func AllTrasactionsId(userid int) Transactions {
	db, err := sql.Open("mysql", dbparams)
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []Transaction{}

	res, err := db.Query("SELECT * FROM `transactions` WHERE `userid` = ?", userid)
	if err != nil {
		log.Fatal("No such user or transactions")
	}

	for res.Next() {
		var tr Transaction
		err = res.Scan(&tr.Id, &tr.Userid, &tr.Email, &tr.Price, &tr.Currency, &tr.CreatedOn, &tr.UpdatedOn, &tr.Status)
		result = append(result, tr)
	}
	defer res.Close()

	return result
}

func AllTransactionsEm(email string) Transactions {
	db, err := sql.Open("mysql", dbparams)
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []Transaction{}

	res, err := db.Query("SELECT * FROM `transactions` WHERE `email` = ?", email)
	if err != nil {
		log.Fatal("No such user or transactions")
	}

	for res.Next() {
		var tr Transaction
		err = res.Scan(&tr.Id, &tr.Userid, &tr.Email, &tr.Price, &tr.Currency, &tr.CreatedOn, &tr.UpdatedOn, &tr.Status)
		result = append(result, tr)
	}
	defer res.Close()

	return result
}
