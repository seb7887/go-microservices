package models

import "github.com/jinzhu/gorm"

type Shipping struct {
	gorm.Model
	OrderId string
	UserId string
	ProductName string
	TotalAmount int32
}