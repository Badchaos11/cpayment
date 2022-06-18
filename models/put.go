package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Reject(t *Transaction) error {
	db, err := sql.Open("mysql", dbparams)
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []TransactionStatus{}
	check, err := db.Query("SELECT `status` FROM `transactions` WHERE `id` = ?", t.Id)
	if err != nil {
		log.Fatal("No such transaction")
	}
	for check.Next() {
		var tr TransactionStatus
		err = check.Scan(&tr.Status)
		result = append(result, tr)
	}

	if result[0].Status == "REJECTED" {
		log.Println("Status cant be changed")
		return err
	} else if result[0].Status == "SUCCESS" {
		log.Println("Status cant be changed")
		return err
	} else if result[0].Status == "UNSUCCESS" {
		log.Println("Status cant be changed")
		return err
	} else {
		res, err := db.Exec("UPDATE transactions SET status = ? WHERE id = ?", t.Status, t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		log.Println("Транзакция успешно отменена")
		return nil
	}
}

func StatusChange(t *Transaction) (int, error) {
	db, err := sql.Open("mysql", dbparams)
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []TransactionStatus{}
	check, err := db.Query("SELECT `status` FROM `transactions` WHERE `id` = ?", t.Id)
	if err != nil {
		log.Fatal("No such transaction")
	}
	for check.Next() {
		var tr TransactionStatus
		err = check.Scan(&tr.Status)
		result = append(result, tr)
	}

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
		res, err := db.Exec("UPDATE `transactions` SET `status` = ? WHERE `id` = ?", t.Status, t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		return 1, nil
	} else {
		res, err := db.Exec("UPDATE `transactions` SET `status` = ? WHERE `id` = ?", t.Status, t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		return 2, nil
	}
}

func StatusChangeWS(t *Transaction) (int, error) {
	db, err := sql.Open("mysql", dbparams)
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []TransactionStatus{}
	check, err := db.Query("SELECT `status` FROM `transactions` WHERE `id` = ?", t.Id)
	if err != nil {
		log.Fatal("No such transaction")
	}
	for check.Next() {
		var tr TransactionStatus
		err = check.Scan(&tr.Status)
		result = append(result, tr)
	}

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
	} else {
		res, err := db.Exec("UPDATE `transactions` SET `status` = ? WHERE `id` = ?", "UNSUCCESS", t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		return 2, nil
	}
}
