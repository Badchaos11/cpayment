package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Структура для обращения к БД
type Transaction struct {
	Id        int     `json:"id"`
	Userid    int     `json:"id_user"`
	Email     string  `json:"email"`
	Price     float32 `json:"price"`
	Currency  string  `json:"currency"`
	CreatedOn string  `json:"created_on"`
	UpdatedOn string  `json:"updated_on"`
	Status    string  `json:"status"`
}

// Струткура для получения статуса записи
type TransactionStatus struct {
	Status string `json:"status"`
}

// Структура для отмены транзакции
type TransactionReject struct {
	Id uint32 `json:"id"`
}

type Transactions []Transaction

// Параметры подключения к БД (временно)
var dbparams string = "badchaos:pe0038900@tcp(127.0.0.1:3306)/constanta"

var db *sql.DB

func Init_DB(params string) error {
	db, err := sql.Open("mysql", params)
	if err != nil {
		log.Println("Failed to open database")
		return err
	}

	err = db.Ping()
	return nil
}
