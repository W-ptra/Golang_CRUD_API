package database

import "time"

type Address struct {
	Street   string
	Province string
	Country  string
}

type Student struct {
	Id        int 		`gorm:"primaryKey;autoIncrement"`
	Name      string
	Age       int
	GPA       float64
	Address   Address
	CreatedAt time.Time
}