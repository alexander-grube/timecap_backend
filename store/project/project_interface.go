package project

import "github.com/spctr-cc/backend-bugtrack/model"

type Store interface {
	GetByID(uint) (*model.Project, error)
	GetAll() ([]*model.Project, error)
	GetAllForAccount(model.Account) ([]*model.Project, error)
	Create(*model.Project) error
	Update(*model.Project) error
	Delete(*model.Project) error
}