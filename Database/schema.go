package database

import (
	"time"
)

type Address struct {
	Id 		 int 		`gorm:"primaryKey;autoIncrement"`
	Street   string
	Province string
	Country  string
}

type Student struct {
	Id        int 		`gorm:"primaryKey;autoIncrement"`
	Name      string
	Age       int
	GPA       float64
	AddressId int
	Address   Address	`gorm:foreignKey:AddressId;referenceId`
	CreatedAt time.Time
}