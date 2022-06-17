package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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

func (t *Transaction) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

type Transactions []*Transaction

func (t *Transactions) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func OneTransaction(id int) TransactionStatus {
	db, err := sql.Open("mysql", "badchaos:pe0038900@tcp(127.0.0.1:3306)/constanta")
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	res, err := db.Query("SELECT `status` FROM `transactions` WHERE `id` = ?", id)
	if err != nil {
		log.Fatal("No such transaction")
	}
	var result TransactionStatus
	err = res.Scan(&result.Status)

	return result
}

func AllTrasactionsId(id int) Transactions {
	db, err := sql.Open("mysql", "badchaos:pe0038900@tcp(127.0.0.1:3306)/constanta")
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []*Transaction{}

	res, err := db.Query("SELECT `status` FROM `transactions` WHERE `id_user` = ?", id)
	if err != nil {
		log.Fatal("No such user or transactions")
	}

	for res.Next() {
		var tr Transaction
		err = res.Scan(&tr.Id, &tr.Userid, &tr.Price, &tr.Currency, &tr.Status)
		result = append(result, &tr)
	}
	defer res.Close()

	return result
}

func AllTransactionsEm(email string) Transactions {
	db, err := sql.Open("mysql", "badchaos:pe0038900@tcp(127.0.0.1:3306)/constanta")
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	result := []*Transaction{}

	res, err := db.Query("SELECT `status` FROM `transactions` WHERE `email` = '?", email)
	if err != nil {
		log.Fatal("No such user or transactions")
	}

	for res.Next() {
		var tr Transaction
		err = res.Scan(&tr.Id, &tr.Email, &tr.Price, &tr.Currency, &tr.Status)
		result = append(result, &tr)
	}
	defer res.Close()

	return result
}

func AddTransaction(t *Transaction) {
	db, err := sql.Open("mysql", "badchaos:pe0038900@tcp(127.0.0.1:3306)/constanta")
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO `transactions` (`userid`, `email`, `price`, `currency`) VALUES (?, ?, ?, ?)", t.Userid, t.Email, t.Price, t.Currency)
	if err != nil {
		log.Fatal("Unable to add user")
	}
	defer insert.Close()
}

func Reject(id int) {
	db, err := sql.Open("mysql", "badchaos:pe0038900@tcp(127.0.0.1:3306)/constanta")
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	res, err := db.Exec("UPDATE transactions SET status = ? WHERE id = ?", "ОТМЕНЕН", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())
}

func StatusChange(t *Transaction) {
	db, err := sql.Open("mysql", "badchaos:pe0038900@tcp(127.0.0.1:3306)/constanta")
	if err != nil {
		log.Fatal("Connection to DB failed")
	}
	defer db.Close()

	res, err := db.Exec("UPDATE transactions SET status = ? WHERE id = ?", t.Status, t.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())
}
