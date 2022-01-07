package types

type Money int64

type Category string

type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

type Account struct {
	ID int64
}
