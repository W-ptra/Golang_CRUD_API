package database

import (
	"time"
)

type Student struct {
	Id        int 		`gorm:"primaryKey;autoIncrement"`
	Name      string
	Age       int
	GPA       float64
	Street    string
	Province  string
	Country   string
	CreatedAt time.Time
}