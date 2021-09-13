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

func (ac *AccountStore) GetByID(id uint) (*model.Account, error) {
	var a model.Account
	if err := ac.db.First(&a, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (ac *AccountStore) Create(a *model.Account) (err error) {
	return ac.db.Create(a).Error
}

func (ac *AccountStore) Update(a *model.Account) (err error) {
	return ac.db.Model(a).Updates(a).Error
}
