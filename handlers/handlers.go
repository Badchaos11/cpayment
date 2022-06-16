package handlers

import (
	"log"
	"net/http"
)

type Transactions struct {
	l *log.Logger
}

func NewTransactions(l *log.Logger) *Transactions {
	return &Transactions{l}
}

func (t *Transactions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
}
