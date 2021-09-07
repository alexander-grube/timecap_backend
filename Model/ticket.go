package model

type Ticket struct {
	Topic     string `json:"topic"`
	Timestamp int64  `json:"timestamp"`
}
