package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountRole uint

const (
	Admin     AccountRole = 1
	Developer AccountRole = 2
	Manager   AccountRole = 3
	User      AccountRole = 4
)

type Account struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string `gorm:"typevarchar(100);uniqueIndex"`
	Password  string
	Role      AccountRole
	Tickets  []Ticket `gorm:"foreignkey:AccountID"`
	Projects []Project `gorm:"many2many:project_accounts;"`
	Organization Organization `gorm:"foreignKey:OrganizationID"`
	OrganizationID uint
}

func (a *Account) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (a *Account) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
	return err == nil
}
