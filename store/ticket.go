package store

import (
	"github.com/spctr-cc/backend-bugtrack/model"
	"gorm.io/gorm"
)

type TicketStore struct {
	db *gorm.DB
}

func NewTicketStore(db *gorm.DB) *TicketStore {
	return &TicketStore{
		db: db,
	}
}

func (ts *TicketStore) GetByID(id uint) (*model.Ticket, error) {
	var a model.Ticket
	if err := ts.db.First(&a, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (ts *TicketStore) Create(a *model.Ticket) (err error) {
	return ts.db.Create(a).Error
}

func (ts *TicketStore) Update(a *model.Ticket) (err error) {
	return ts.db.Model(a).Updates(a).Error
}

func (ts *TicketStore) Delete(a *model.Ticket) (err error) {
	return ts.db.Delete(a).Error
}