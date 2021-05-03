package entity

type Account struct {
	initialized    bool
	activeCard     bool
	availableLimit Amount
}

func NewAccountEmpty() Account {
	return Account{
		initialized:    false,
		activeCard:     false,
		availableLimit: NewAmount(0),
	}
}

func NewAccount(activeCard bool, availableLimit uint) Account {
	return Account{
		initialized:    true,
		activeCard:     activeCard,
		availableLimit: NewAmount(availableLimit),
	}
}

func NewAccountSubtractLimit(account Account, transaction Transaction) Account {
	return Account{
		initialized:    account.initialized,
		activeCard:     account.activeCard,
		availableLimit: NewAmount(account.availableLimit.value - transaction.amount.value),
	}
}

func (a Account) IsInitialized() bool {
	return a.initialized
}

func (a Account) IsActiveCard() bool {
	return a.activeCard
}

func (a Account) GetAvailableLimit() Amount {
	return a.availableLimit
}
