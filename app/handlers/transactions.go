package handlers

import (
	"log"
)

type Transactions struct {
	l *log.Logger
}

var token string = "4hbkjdznfk3i27ecb1"

func NewTransactions(l *log.Logger) *Transactions {
	return &Transactions{l}
}
