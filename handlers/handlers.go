package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Badchaos11/cpayment/models"
)

type Transactions struct {
	l *log.Logger
}

func NewTransactions(l *log.Logger) *Transactions {
	return &Transactions{l}
}

func (t *Transactions) GetOneById(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Transaction by ID")

	ot := models.OneTransaction(1)
	fmt.Println(ot)
}

func (t *Transactions) GetAllById(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET All Transactions by User ID")
	return
}

func (t *Transactions) GetAllByEmail(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET All Transaction by User EMail")
	return
}

func (t *Transactions) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Create new transaction")
	return
}

func (t *Transactions) RejectTransaction(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PATCH reject transaction by ID")
	return
}

func (t *Transactions) ChangeTransactionStatus(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PATCH Change transaction status by System")
	return
}

type KeyTransaction struct{}
