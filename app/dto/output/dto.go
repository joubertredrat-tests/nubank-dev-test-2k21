package output

type AccountLine struct {
	Account struct {
		ActiveCard     bool     `json:"active-card"`
		AvailableLimit uint     `json:"available-limit"`
		Violations     []string `json:"violations"`
	} `json:"account"`
}

func NewAccountLine(activeCard bool, availableLimit uint, violations []string) AccountLine {
	return AccountLine{
		Account: struct {
			ActiveCard     bool     `json:"active-card"`
			AvailableLimit uint     `json:"available-limit"`
			Violations     []string `json:"violations"`
		}{
			ActiveCard:     activeCard,
			AvailableLimit: availableLimit,
			Violations:     violations,
		},
	}
}
