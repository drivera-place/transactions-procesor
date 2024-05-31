package domain

import (
	"time"
)

type Transaction struct {
	id   int       `json:Id`
	date time.Time `json:Date`
	txn  float64   `json:transaction`
}

func (t Transaction) New(id int, date time.Time, txn float64) *Transaction {
	return &Transaction{id: id, date: date, txn: txn}
}

func (t *Transaction) Id() int {
	return t.id
}

func (t *Transaction) Date() time.Time {
	return t.date
}

func (t *Transaction) Txn() float64 {
	return t.txn
}
