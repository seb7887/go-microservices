package models

import "github.com/jinzhu/gorm"

type Login struct {
	gorm.Model
	Hash string
	Email string
}

