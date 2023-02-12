package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name 	  string `gorm:"typevarchar(100);uniqueIndex"`
	Accounts  []Account `gorm:"foreignkey:OrganizationID"`
}