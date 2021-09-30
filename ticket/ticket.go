package ticket

import (
	"github.com/spctr-cc/backend-bugtrack/model"
)

type Store interface {
	GetByID(uint) (*model.Ticket, error)
	Create(*model.Ticket) error
	Update(*model.Ticket) error
}
