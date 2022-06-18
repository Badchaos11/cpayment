package handlers

import (
	"fmt"
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
	res, err := models.Reject(&tr)
	if err != nil {
		t.l.Println("Что-то пошло не так")
	}
	if res == true {
		t.l.Println("Транзакция успешно отменена")
		fmt.Fprint(w, "Транзакция успешно отменена")
	}
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
		t.l.Println("Статус изменить невозможно")
		fmt.Fprint(w, "Статус изменить невозможно")
	}
	if res == 1 {
		t.l.Println("Статус транзакции успешено установлен: SUCCESS")
		fmt.Fprint(w, "Статус транзакции успешено установлен: SUCCESS")
	} else if res == 2 {
		t.l.Println("Статус транзакции успешно установлен: UNSUCCESS")
		fmt.Fprint(w, "Статус транзакции успешно установлен: UNSUCCESS")
	} else {
		t.l.Println("Статус изменить невозможно")
		fmt.Fprint(w, "Статус изменить невозможно")
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
		t.l.Println("Статус изменить невозможно")
		fmt.Fprint(w, "Статус изменить невозможно")
	}
	if res == 1 {
		t.l.Println("Статус транзакции установлен: SUCCESS")
		fmt.Fprint(w, "Статус транзакции установлен: SUCCESS")
	} else if res == 2 {
		t.l.Println("Статус транзакции установлен: UNSUCCESS")
		fmt.Fprint(w, "Статус транзакции установлен: UNSUCCESS")
	} else {
		t.l.Println("Статус изменить невозможно")
		fmt.Fprint(w, "Статус изменить невозможно")
	}

}
