package store

import (
	"github.com/spctr-cc/backend-bugtrack/model"
	"gorm.io/gorm"
)

type ProjectStore struct {
	db *gorm.DB
}

func NewProjectStore(db *gorm.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

func (ps *ProjectStore) GetByID(id uint) (*model.Project, error) {
	var p model.Project
	if err := ps.db.First(&p, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &p, nil
}

func (ps *ProjectStore) GetAll() ([]*model.Project, error) {
	var p []*model.Project
	if err := ps.db.Find(&p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

func (ps *ProjectStore) GetAllForAccount(account model.Account) ([]*model.Project, error) {
	var p []*model.Project
	if err := ps.db.Model(&account).Association("Projects").Find(&p); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

func (ps *ProjectStore) Create(p *model.Project) (err error) {
	return ps.db.Create(p).Error
}

func (ps *ProjectStore) Update(p *model.Project) (err error) {
	return ps.db.Model(p).Updates(p).Error
}

func (ps *ProjectStore) Delete(p *model.Project) (err error) {
	return ps.db.Delete(p).Error
}
