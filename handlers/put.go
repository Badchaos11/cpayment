package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Badchaos11/cpayment/models"
)

// Вызов функции отмены транзакции по PUT запросу
func (t *Transactions) RejectTransaction(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PUT reject transaction by ID")
	d, err := ioutil.ReadAll(r.Body) // чтение тела запроса
	if err != nil {
		log.Fatal("Panic")
	}
	t.l.Println(d)
	st := "REJECTED"
	sid := string(d)             // получение id транзакции
	id, err := strconv.Atoi(sid) // преобразование id в int
	if err != nil {
		log.Fatal("Panic")
	}
	tr := models.Transaction{Id: id, Status: st} // формирование структуры для вызова БД
	res, err := models.Reject(&tr)               // запрос к БД
	if err != nil {
		t.l.Println("Что-то пошло не так")
	}
	if res == true { // сообщение об успехе операции
		t.l.Println("Транзакция успешно отменена")   // вывод в консоль
		fmt.Fprint(w, "Транзакция успешно отменена") // ответ пользователю
	}
}

// Вызов функции изменения статуса платежа платежной системой PUT запрос
// Доступно только по токену авторизации в заголовке
func (t *Transactions) ChangeTransactionStatus(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle PUT Change transaction status by System")
	d, err := ioutil.ReadAll(r.Body) // чтение тела запроса
	if err != nil {
		log.Fatal("Panic")
	}
	t.l.Println(d)
	sid := string(d)             // получение id транзакции
	id, err := strconv.Atoi(sid) // преобразование id в int
	if err != nil {
		log.Fatal("Panic")
	}

	tr := models.Transaction{Id: id} // формирование структуры для вызова БД

	res, err := models.StatusChangeWS(&tr) // запрос к БД
	if err != nil {                        // сообщение о невозможности изменить статус
		t.l.Println("Статус изменить невозможно")
		fmt.Fprint(w, "Статус изменить невозможно")
	}
	if res == 1 { // сообщение об успехе
		t.l.Println("Статус транзакции успешено установлен: SUCCESS")
		fmt.Fprint(w, "Статус транзакции успешено установлен: SUCCESS")
	} else if res == 2 { // сообщение об успехе
		t.l.Println("Статус транзакции успешно установлен: UNSUCCESS")
		fmt.Fprint(w, "Статус транзакции успешно установлен: UNSUCCESS")
	} else {
		t.l.Println("Статус изменить невозможно")
		fmt.Fprint(w, "Статус изменить невозможно")
	}
}
