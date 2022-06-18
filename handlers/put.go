package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Badchaos11/cpayment/models"
)

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
	if res == 1 {
		t.l.Println("Статус транзакции успешено установлен: SUCCESS")
	} else if res == 2 {
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
	if res == 1 {
		t.l.Println("Статус транзакции установлен: SUCCESS")
	} else if res == 2 {
		t.l.Println("Статус транзакции установлен: UNSUCCESS")
	}
}
