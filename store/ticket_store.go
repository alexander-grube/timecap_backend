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
	var t model.Ticket
	if err := ts.db.First(&t, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &t, nil
}

func (ts *TicketStore) GetAll() ([]*model.Ticket, error) {
	var t []*model.Ticket
	if err := ts.db.Find(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return t, nil
}

func (ts *TicketStore) Create(t *model.Ticket) (err error) {
	return ts.db.Create(t).Error
}

func (ts *TicketStore) Update(t *model.Ticket) (err error) {
	return ts.db.Model(t).Updates(t).Error
}

func (ts *TicketStore) Delete(t *model.Ticket) (err error) {
	return ts.db.Delete(t).Error
}