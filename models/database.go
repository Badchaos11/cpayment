package models

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"time"
)

type Transaction struct {
	Id_transaction int       `json:"id"`
	Id_user        int       `json:"id_user"`
	Email_user     string    `json:"email"`
	Price          float32   `json:"price"`
	Currency       string    `json:"currency"`
	CreatedOn      time.Time `json:"-"`
	UpdatedOn      time.Time `json:"t-"`
	Status         string    `json:"-"`
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

func OneTransaction(id int) Transaction {
	db, err := sql.Open("postgres", "host=localhost dbname=Test sslmode=disable user=postgres password=secret")

	if err != nil {
		log.Fatal("Connection to DB failed")
	}

	defer db.Close()

	one_transaction := Transaction{}

	return one_transaction
}

func AllTrasactionsId(id int) Transactions {
	db, err := sql.Open("postgres", "host=localhost dbname=Test sslmode=disable user=postgres password=secret")

	if err != nil {
		log.Fatal("Connection to DB failed")
	}

	defer db.Close()

	result := []*Transaction{
		&Transaction{},
		&Transaction{},
	}

	return result
}

func AllTransactionsEm(email string) Transactions {
	db, err := sql.Open("postgres", "host=localhost dbname=Test sslmode=disable user=postgres password=secret")

	if err != nil {
		log.Fatal("Connection to DB failed")
	}

	defer db.Close()

	result := []*Transaction{
		&Transaction{},
		&Transaction{},
	}

	return result
}

func AddTransaction(t *Transaction) {
	db, err := sql.Open("postgres", "host=localhost dbname=Test sslmode=disable user=postgres password=secret")

	if err != nil {
		log.Fatal("Connection to DB failed")
	}

	defer db.Close()

}

func Reject(id int) {
	db, err := sql.Open("postgres", "host=localhost dbname=Test sslmode=disable user=postgres password=secret")

	if err != nil {
		log.Fatal("Connection to DB failed")
	}

	defer db.Close()

}

func Status(id int, status string) {
	db, err := sql.Open("postgres", "host=localhost dbname=Test sslmode=disable user=postgres password=secret")

	if err != nil {
		log.Fatal("Connection to DB failed")
	}

	defer db.Close()
}
