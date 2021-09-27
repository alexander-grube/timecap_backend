package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountRole int

const (
	Admin     AccountRole = 0
	Developer AccountRole = 1
	Manager   AccountRole = 2
	User      AccountRole = 3
)

type Account struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string `gorm:"typevarchar(100);uniqueIndex"`
	Password  string
	Role      AccountRole
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
