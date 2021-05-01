package entity

import (
	"time"
)

type Transaction struct {
	merchant string
	amount   Amount
	time     time.Time
}

func NewTransaction(merchant string, amount uint, time time.Time) Transaction {
	return Transaction{
		merchant: merchant,
		amount:   NewAmount(amount),
		time:     time,
	}
}

func (t Transaction) GetMerchant() string {
	return t.merchant
}

func (t Transaction) GetAmount() Amount {
	return t.amount
}

func (t Transaction) GetTime() time.Time {
	return t.time
}
