package account

import (
	"github.com/spctr-cc/backend-bugtrack/model"
)

type Store interface {
	GetByID(uint) (*model.Account, error)
	Create(*model.Account) error
	Update(*model.Account) error
}
