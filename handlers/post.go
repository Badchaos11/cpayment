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

// Вызов функции создания новой транзакции POST запрос
func (t *Transactions) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Create new transaction")
	d, err := ioutil.ReadAll(r.Body) // получение данных из тела запроса
	if err != nil {
		log.Fatal("Panic")
	}
	rs := strings.Split(string(d), "&") // разделение данных на отдельные строки
	uid, err := strconv.Atoi(rs[0])     // преобразование userid в int
	if err != nil {
		log.Fatal("Invalid ID")
	}
	em := rs[1]                            // получение email
	p, err := strconv.ParseFloat(rs[2], 4) // получение суммы платежа, перевод во float
	if err != nil {
		log.Fatal("Invalid Amount of Money")
	}
	cr := rs[3] // получение типа валюты
	pr := float32(p)
	tr := models.Transaction{Userid: uid, Email: em, Price: pr, Currency: cr} // формирование структуры для вызова  БД
	t.l.Println(tr)
	res := models.AddTransaction(&tr) // запрос к БД
	if res == true {
		fmt.Fprint(w, "Транзакция успешно создана") // сообщение об успешном выполнении
	}
}

// Вызов функции создания новой транзакции POST запрос
// Создается с ошибкой
func (t *Transactions) CreateTransactionFail(w http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Create new transaction")
	d, err := ioutil.ReadAll(r.Body) // получение данных из тела запроса
	if err != nil {
		log.Fatal("Panic")
	}
	rs := strings.Split(string(d), "&") // разделение данных на отдельные строки
	uid, err := strconv.Atoi(rs[0])     // преобразование userid в int
	if err != nil {
		log.Fatal("Invalid ID")
	}
	em := rs[1]                            // получение email
	p, err := strconv.ParseFloat(rs[2], 4) // получение суммы платежа, перевод во float
	if err != nil {
		log.Fatal("Invalid Amount of Money")
	}
	cr := rs[3] // получение типа валюты
	pr := float32(p)
	st := "FAILED"
	tr := models.Transaction{Userid: uid, Email: em, Price: pr, Currency: cr, Status: st} // формирование структуры для вызова  БД
	t.l.Println(tr)
	res := models.AddTransactionFail(&tr) // запрос к БД
	if res == true {
		fmt.Fprint(w, "При добавлении транзакции произошла ошибка") // сообщение об успешном выполнении
	}
}
