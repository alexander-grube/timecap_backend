package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `gorm:"typevarchar(100);uniqueIndex"`
	Description string
	Tickets     []Ticket  `gorm:"foreignKey:ProjectID"`
	Accounts    []Account `gorm:"many2many:project_accounts;"`
}
