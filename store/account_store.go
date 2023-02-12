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
			return nil, err
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

func (as *AccountStore) AddToProject(a *model.Account, p *model.Project) error {
	return as.db.Model(a).Association("Projects").Append(p)
}

func (as *AccountStore) RemoveFromProject(a *model.Account, p *model.Project) error {
	return as.db.Model(a).Association("Projects").Delete(p)
}

func (as *AccountStore) AddToOrganization(a *model.Account, o *model.Organization) error {
	return as.db.Model(a).Association("Organizations").Append(o)
}

func (as *AccountStore) RemoveFromOrganization(a *model.Account, o *model.Organization) error {
	return as.db.Model(a).Association("Organizations").Delete(o)
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
