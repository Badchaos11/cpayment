package handlers

import (
	"log"
)

// Структура для создания логера
type Transactions struct {
	l *log.Logger
}

// Токен для авторизации (временно)
var token string = "4hbkjdznfk3i27ecb1"

// Создание нового логера
func NewTransactions(l *log.Logger) *Transactions {
	return &Transactions{l}
}
