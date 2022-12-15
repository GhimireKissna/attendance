package model

type User struct {
	ID       	int `gorm:"primary_key"`
	FirstName   string
	MiddleName  string
	LastName    string
	Designation string
}