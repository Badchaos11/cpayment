package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Badchaos11/cpayment/models"
	"github.com/gorilla/mux"
)

// Функция для получения статуса транзакции по её ID
func (t *Transactions) GetOneById(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET Transaction by ID")
	vars := mux.Vars(r) // получение ID
	t.l.Println(vars)
	id, err := strconv.Atoi(vars["id"]) // конвертация в int
	if err != nil {
		log.Fatal("Invalid ID entered")
	}
	ot := models.OneTransaction(id)                    // запрос к БД
	t.l.Println("Статус транзакции: ", ot[0].Status)   // вывод результата в консоль
	fmt.Fprint(w, "Статус транзакции: ", ot[0].Status) // вывод результата в ответ на запрос
}

// Функция получения всех транзакций пользователя по его ID
func (t *Transactions) GetAllById(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET All Transactions by User ID")
	vars := mux.Vars(r) // получение ID пользователя
	t.l.Println(vars)
	id, err := strconv.Atoi(vars["userid"]) // конвертация в int
	if err != nil {
		log.Fatal("Invalid User ID entered")
	}
	trs := models.AllTrasactionsId(id) // запрос к БД
	for i := 0; i < len(trs); i++ {    // вывод всех результатов в цикле
		t.l.Println(trs[i])                                  // вывод результата в консоль
		fmt.Fprint(w, "Данные транзакции: ", trs[i], "\r\n") // вывод результата в ответ на запрос
	}
}

// Функция получения всех транзакций пользователя по его EMail
func (t *Transactions) GetAllByEmail(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle GET All Transaction by User EMail")
	vars := mux.Vars(r) // получение email
	t.l.Println(vars)
	email := vars["email"]
	trs := models.AllTransactionsEm(email) // запрос к БД
	for i := 0; i < len(trs); i++ {        // вывод всех результатов в цикле
		t.l.Println(trs[i])                                 // вывод результата в консоль
		fmt.Fprint(w, "Данные транзакции: ", trs[i], " \n") // вывод результата в ответ на запрос
	}
}
