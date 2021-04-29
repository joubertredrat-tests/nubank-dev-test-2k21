package command

import (
	"time"
)

const (
	LINE_ACCOUNT     = "account"
	LINE_TRANSACTION = "transaction"
)

type OperationLine interface {
	GetType() string
}

type AccountLine struct {
	Account struct {
		ActiveCard     bool `json:"active-card"`
		AvailableLimit uint `json:"available-limit"`
	} `json:"account"`
}

func (a AccountLine) GetType() string {
	return LINE_ACCOUNT
}

type TransactionLine struct {
	Transaction struct {
		Merchant string    `json:"merchant"`
		Amount   uint      `json:"amount"`
		Time     time.Time `json:"time"`
	} `json:"transaction"`
}

func (t TransactionLine) GetType() string {
	return LINE_TRANSACTION
}

type Operations struct {
	Lines []OperationLine
}

func (o *Operations) AddLine(l OperationLine) {
	o.Lines = append(o.Lines, l)
}

func NewOperations() Operations {
	return Operations{
		Lines: []OperationLine{},
	}
}
