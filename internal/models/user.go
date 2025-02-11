package models

type User struct {
	ID       int
	Username string
	Password string

	Balance         int
	Inventory       []Item
	TransferHistory History
}
