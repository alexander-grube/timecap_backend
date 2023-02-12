package model

import "gorm.io/gorm"

type TicketStatus uint

const (
	New        TicketStatus = 1
	Open       TicketStatus = 2
	InProgress TicketStatus = 3
	Resolved   TicketStatus = 4
)

type TicketType uint

const (
	Other   TicketType = 1
	Bug     TicketType = 2
	Feature TicketType = 3
)

type TicketPriority uint

const (
	Low    TicketPriority = 1
	Medium TicketPriority = 2
	High   TicketPriority = 3
)

type Ticket struct {
	gorm.Model
	Topic       string
	Description string
	Priority    TicketPriority
	Type        TicketType
	Status      TicketStatus
	AccountID   uint
	ProjectID   uint
}
