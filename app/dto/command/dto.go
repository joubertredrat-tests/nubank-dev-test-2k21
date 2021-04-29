package command

import (
	"fmt"
	"time"
)

const (
	LINE_ACCOUNT     = "account"
	LINE_TRANSACTION = "transaction"
)

type Line struct {
	Type string `json:""`
}

func (l *Line) UnmarshalJSON(b []byte) error {
	fmt.Println(b)
	return nil
}

type AccountLine struct {
	Account struct {
		ActiveCard     bool `json:"active-card"`
		AvailableLimit uint `json:"available-limit"`
	} `json:"account"`
}

type TransactionLine struct {
	Transaction struct {
		Merchant string    `json:"merchant"`
		Amount   int       `json:"amount"`
		Time     time.Time `json:"time"`
	} `json:"transaction"`
}
