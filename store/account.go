package store

import (
	"github.com/spctr-cc/backend-bugtrack/model"
	"gorm.io/gorm"
)

type AccountStore struct {
	db *gorm.DB
}

func NewAccountStore(db *gorm.DB) *AccountStore {
	return &AccountStore{
		db: db,
	}
}

func (as *AccountStore) GetByID(id uint) (*model.Account, error) {
	var a model.Account
	if err := as.db.First(&a, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (as *AccountStore) GetByEmail(e string) (*model.Account, error) {
	var a model.Account
	if err := as.db.Where(&model.Account{Email: e}).First(&a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (as *AccountStore) Create(a *model.Account) (err error) {
	return as.db.Create(a).Error
}

func (as *AccountStore) Update(a *model.Account) (err error) {
	return as.db.Model(a).Updates(a).Error
}

func (as *AccountStore) Delete(a *model.Account) (err error) {
	return as.db.Delete(a).Error
}
