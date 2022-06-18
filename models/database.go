package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Transaction struct {
	Id        int     `json:"id"`
	Userid    int     `json:"id_user"`
	Email     string  `json:"email"`
	Price     float32 `json:"price"`
	Currency  string  `json:"currency"`
	CreatedOn string  `json:"created_on"`
	UpdatedOn string  `json:"updated_on"`
	Status    string  `json:"status"`
}

type TransactionStatus struct {
	Status string `json:"status"`
}

type TransactionReject struct {
	Id uint32 `json:"id"`
}

type Transactions []Transaction

var dbparams string = "badchaos:pe0038900@tcp(127.0.0.1:3306)/constanta"

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

func StatusChange(t *Transaction) (bool, error) {
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
		return false, err
	} else if result[0].Status == "SUCCESS" {
		log.Println("Status cant be changed")
		return false, err
	} else if result[0].Status == "UNSUCCESS" {
		log.Println("Status cant be changed")
		return false, err
	} else if result[0].Status == "NEW" {
		res, err := db.Exec("UPDATE `transactions` SET `status` = ? WHERE `id` = ?", t.Status, t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		return true, nil
	} else {
		res, err := db.Exec("UPDATE `transactions` SET `status` = ? WHERE `id` = ?", t.Status, t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		return false, nil
	}
}

func StatusChangeWS(t *Transaction) (bool, error) {
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
		return false, err
	} else if result[0].Status == "SUCCESS" {
		log.Println("Status cant be changed")
		return false, err
	} else if result[0].Status == "UNSUCCESS" {
		log.Println("Status cant be changed")
		return false, err
	} else if result[0].Status == "NEW" {
		res, err := db.Exec("UPDATE `transactions` SET `status` = ? WHERE `id` = ?", "SUCCES", t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		return true, nil
	} else {
		res, err := db.Exec("UPDATE `transactions` SET `status` = ? WHERE `id` = ?", "UNSUCCESS", t.Id)
		if err != nil {
			panic(err)
		}
		log.Println(res.RowsAffected())
		return false, nil
	}
}
