package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	t.l.Println(vars)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Invalid ID entered")
	}
	ot := models.OneTransaction(id)
	t.l.Println(ot)
}

func (t *Transactions) GetAllById(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET All Transactions by User ID")
	vars := mux.Vars(r)
	t.l.Println(vars)
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
	t.l.Println(vars)
	email := vars["email"]
	trs := models.AllTransactionsEm(email)
	fmt.Println(trs)
	return
}

func (t *Transactions) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Create new transaction")
	tr := r.Context().Value(KeyTransaction{}).(models.Transaction)
	t.l.Println(tr)

	models.AddTransaction(&tr)
	return
}

func (t *Transactions) RejectTransaction(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PATCH reject transaction by ID")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Panic")
	}

	t.l.Println(d[0])
	models.Reject(int(d[0]))

}

func (t *Transactions) ChangeTransactionStatus(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PATCH Change transaction status by System")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Panic")
	}

	rs := strings.Split(string(d), "&")
	st := rs[0]
	id, err := strconv.Atoi(rs[1])
	if err != nil {
		log.Fatal("Invalid ID")
	}

	tr := models.Transaction{Id: id, Status: st}

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
