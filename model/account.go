package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Firstname string `json:"firstname" validate:"required,min=2,max=100"`
	Lastname  string `json:"lastname" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email" gorm:"typevarchar(100);uniqueIndex"`
	Password  string `json:"password" validate:"required"`
	Mobile    string `json:"mobile" validate:"required"`
	Street    string `json:"street" validate:"required,min=2,max=100"`
	Zipcode   string `json:"zipcode" validate:"required"`
	City      string `json:"city" validate:"required"`
	Country   string `json:"country" validate:"required"`
}
