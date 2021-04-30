package input

import (
	"time"
)

const (
	LINE_ACCOUNT     = "account"
	LINE_TRANSACTION = "transaction"
)

type OperationLine interface {
	IsAccount() bool
	IsTransaction() bool
}

type AccountLine struct {
	Account struct {
		ActiveCard     bool `json:"active-card"`
		AvailableLimit uint `json:"available-limit"`
	} `json:"account"`
}

func (a AccountLine) IsAccount() bool {
	return true
}

func (a AccountLine) IsTransaction() bool {
	return false
}

type TransactionLine struct {
	Transaction struct {
		Merchant string    `json:"merchant"`
		Amount   uint      `json:"amount"`
		Time     time.Time `json:"time"`
	} `json:"transaction"`
}

func (t TransactionLine) IsAccount() bool {
	return false
}

func (t TransactionLine) IsTransaction() bool {
	return true
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
