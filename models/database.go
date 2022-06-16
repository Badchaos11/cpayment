package models

import (
	"encoding/json"
	"io"
	"time"
)

type Transaction struct {
	Id_transaction int       `json:"id"`
	Id_user        int       `json:"id_user"`
	Email_user     string    `json:"email"`
	Price          float32   `json:"price"`
	Currency       string    `json:"currency"`
	CreatedOn      time.Time `json:"time_creation"`
	UpdatedOn      time.Time `json:"time_updated"`
	Status         string    `json:"status"`
}

func (t *Transaction) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

type Transactions []*Transaction

func (t *Transactions) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}
