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
	res := models.AddTransaction(&tr)
	if res == true {
		fmt.Fprint(w, "Транзакция успешно создана")
	}
}

func (t *Transactions) CreateTransactionFail(w http.ResponseWriter, r *http.Request) {
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
	st := "FAILED"
	tr := models.Transaction{Userid: uid, Email: em, Price: pr, Currency: cr, Status: st}
	t.l.Println(tr)
	res := models.AddTransactionFail(&tr)
	if res == true {
		fmt.Fprint(w, "При добавлении транзакции произошла ошибка")
	}
}
