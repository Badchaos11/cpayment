package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Отмена транзакции
func Reject(t *Transaction) (bool, error) {
	db, err := sql.Open("mysql", dbparams) // подключение к БД
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []TransactionStatus{}
	check, err := db.Query("SELECT `status` FROM `transactions` WHERE `id` = ?", t.Id) // запрос к БД, проверка статуса транзакции
	if err != nil {
		log.Fatal("No such transaction")
	}
	for check.Next() {
		var tr TransactionStatus
		err = check.Scan(&tr.Status)
		result = append(result, tr)
	}
	// Сообщения о невозможности изменения статуса
	if result[0].Status == "REJECTED" {
		log.Println("Status cant be changed")
		return false, err
	} else if result[0].Status == "SUCCESS" {
		log.Println("Status cant be changed")
		return false, err
	} else if result[0].Status == "UNSUCCESS" {
		log.Println("Status cant be changed")
		return false, err
	} else { // изменение записи в БД, сообщение об успехе
		res, err := db.Exec("UPDATE transactions SET status = ? WHERE id = ?", t.Status, t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		log.Println("Транзакция успешно отменена")
		return true, nil
	}
}

// Изменение статуса платежа платежной системой.
func StatusChangeWS(t *Transaction) (int, error) {
	db, err := sql.Open("mysql", dbparams) // подключение к БД
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []TransactionStatus{}
	check, err := db.Query("SELECT `status` FROM `transactions` WHERE `id` = ?", t.Id) // запрос к БД, проверка статуса транзакции
	if err != nil {
		log.Fatal("No such transaction")
	}
	for check.Next() {
		var tr TransactionStatus
		err = check.Scan(&tr.Status)
		result = append(result, tr)
	}
	// Сообщения о невозможности изменения статуса
	if result[0].Status == "REJECTED" {
		log.Println("Status cant be changed")
		return 0, err
	} else if result[0].Status == "SUCCESS" {
		log.Println("Status cant be changed")
		return 0, err
	} else if result[0].Status == "UNSUCCESS" {
		log.Println("Status cant be changed")
		return 0, err
	} else if result[0].Status == "NEW" {
		res, err := db.Exec("UPDATE `transactions` SET `status` = ? WHERE `id` = ?", "SUCCES", t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		return 1, nil
	} else { // изменение записи в БД, сообщение об успехе
		res, err := db.Exec("UPDATE `transactions` SET `status` = ? WHERE `id` = ?", "UNSUCCESS", t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		return 2, nil
	}
}
