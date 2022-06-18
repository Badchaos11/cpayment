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
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Panic")
	}
	rs := strings.Split(string(d), "&")
	uid, err := strconv.Atoi(rs[0])
	if err != nil {
		log.Fatal("Invalid ID")
	}
	em := rs[1]
	p, err := strconv.ParseFloat(rs[2], 4)
	if err != nil {
		log.Fatal("Invalid Amount of Money")
	}
	cr := rs[3]
	pr := float32(p)
	tr := models.Transaction{Userid: uid, Email: em, Price: pr, Currency: cr}
	t.l.Println(tr)
	models.AddTransaction(&tr)
	return
}

func (t *Transactions) RejectTransaction(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PUT reject transaction by ID")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Panic")
	}
	t.l.Println(d)
	st := "REJECTED"
	sid := string(d)
	id, err := strconv.Atoi(sid)
	if err != nil {
		log.Fatal("Panic")
	}
	tr := models.Transaction{Id: id, Status: st}
	models.Reject(&tr)
}

func (t *Transactions) ChangeTransactionStatusWS(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PUT Change transaction status by System")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Panic")
	}
	t.l.Println(d)
	sid := string(d)
	id, err := strconv.Atoi(sid)
	if err != nil {
		log.Fatal("Panic")
	}

	tr := models.Transaction{Id: id}

	res, err := models.StatusChangeWS(&tr)
	if err != nil {
		t.l.Println("Status was not changed")
	}
	if res == true {
		t.l.Println("Статус транзакции успешено установлен: SUCCESS")
	} else {
		t.l.Println("Статус транзакции успешно установлен: UNSUCCESS")
	}
}

func (t *Transactions) ChangeTransactionStatus(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PUT Change transaction status by System")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Panic")
	}
	t.l.Println(d)
	rs := strings.Split(string(d), "&")
	st := rs[0]
	id, err := strconv.Atoi(rs[1])
	if err != nil {
		log.Fatal("Invalid ID")
	}

	tr := models.Transaction{Id: id, Status: st}

	res, err := models.StatusChange(&tr)
	if err != nil {
		t.l.Println("Status was not changed")
	}
	if res == true {
		t.l.Println("Статус транзакции установлен: SUCCESS")
	} else {
		t.l.Println("Статус транзакции установлен: UNSUCCESS")
	}
}

func (t *Transactions) MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header["Authorization"][0]
		t.l.Println(h)
		if h != token {
			log.Fatal("Invalid authorization header received")
			return
		}
		t.l.Println("Token allowed, go to next handler")
		next.ServeHTTP(w, r)
	})
}
