package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Badchaos11/cpayment/models"
	"github.com/gorilla/mux"
)

func (t *Transactions) GetOneById(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Transaction by ID")
	vars := mux.Vars(r)
	t.l.Println(vars)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Invalid ID entered")
	}
	ot := models.OneTransaction(id)
	t.l.Println("Статус транзакции: ", ot[0].Status)
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
	for i := 0; i < len(trs); i++ {
		t.l.Println(trs[i])
	}
}

func (t *Transactions) GetAllByEmail(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET All Transaction by User EMail")
	vars := mux.Vars(r)
	t.l.Println(vars)
	email := vars["email"]
	trs := models.AllTransactionsEm(email)
	for i := 0; i < len(trs); i++ {
		t.l.Println(trs[i])
	}
}
