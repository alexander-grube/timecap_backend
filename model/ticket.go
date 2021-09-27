package model

import "gorm.io/gorm"

type TicketStatus int

const (
	New        TicketStatus = 0
	Open       TicketStatus = 1
	InProgress TicketStatus = 2
	Resolved   TicketStatus = 3
)

type TicketType int

const (
	Other   TicketType = 0
	Bug     TicketType = 1
	Feature TicketType = 2
)

type TicketPriority int

const (
	None   TicketPriority = 0
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
	Timestamp   int64
	AccountID   int
}
