package models

import (
	_ "github.com/go-sql-driver/mysql"
)

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

type TransactionStatus struct {
	Status string `json:"status"`
}

type TransactionReject struct {
	Id uint32 `json:"id"`
}

type Transactions []Transaction

var dbparams string = "badchaos:pe0038900@tcp(127.0.0.1:3306)/constanta"
