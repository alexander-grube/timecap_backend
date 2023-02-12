package account

import (
	"github.com/spctr-cc/backend-bugtrack/model"
)

type Store interface {
	GetByID(uint) (*model.Account, error)
	GetByEmail(string) (*model.Account, error)
	AddToProject(*model.Account, *model.Project) error
	RemoveFromProject(*model.Account, *model.Project) error
	AddToOrganization(*model.Account, *model.Organization) error
	RemoveFromOrganization(*model.Account, *model.Organization) error
	Create(*model.Account) error
	Update(*model.Account) error
	Delete(*model.Account) error
}
