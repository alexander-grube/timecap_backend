package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Topic     string `json:"topic"`
	Timestamp int64  `json:"timestamp"`
	AccountID int
}
