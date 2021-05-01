package entity

type Amount struct {
	value uint
}

func NewAmount(value uint) Amount {
	return Amount{
		value: value,
	}
}

func (a Amount) GetValue() uint {
	return a.value
}
