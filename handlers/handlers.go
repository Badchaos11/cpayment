package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Badchaos11/cpayment/models"
	"github.com/gorilla/mux"
)

type Transactions struct {
	l *log.Logger
}

var token string = "4hbkjdznfk3i27ecb1"

func NewTransactions(l *log.Logger) *Transactions {
	return &Transactions{l}
}

func (t *Transactions) GetOneById(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Transaction by ID")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Invalid ID entered")
	}
	ot := models.OneTransaction(id)
	fmt.Println(ot)
}

func (t *Transactions) GetAllById(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET All Transactions by User ID")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["userid"])
	if err != nil {
		log.Fatal("Invalid User ID entered")
	}
	trs := models.AllTrasactionsId(id)
	fmt.Println(trs)
	return
}

func (t *Transactions) GetAllByEmail(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET All Transaction by User EMail")
	vars := mux.Vars(r)
	email := vars["email"]

	trs := models.AllTransactionsEm(email)
	fmt.Println(trs)
	return
}

func (t *Transactions) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Create new transaction")
	tr := r.Context().Value(KeyTransaction{}).(models.Transaction)

	models.AddTransaction(&tr)
	return
}

func (t *Transactions) RejectTransaction(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PATCH reject transaction by ID")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["userid"])
	if err != nil {
		log.Fatal("Invalid User ID entered")
	}
	models.Reject(id)

}

func (t *Transactions) ChangeTransactionStatus(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PATCH Change transaction status by System")
	tr := r.Context().Value(KeyTransaction{}).(models.Transaction)

	models.StatusChange(&tr)
}

type KeyTransaction struct{}

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header["Authorization"][1]
		fmt.Println(h)
		if h != token {
			log.Fatal("Invalid authorization header received")
			return
		}
		log.Println("Token allowed, go to next handler")
		next.ServeHTTP(w, r)
	})
}
