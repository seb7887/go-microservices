package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	UserId string
	ProductName string
	TotalAmount int32
}