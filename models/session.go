package models

import "time"

type Session struct {
	ID             int
	UserID         int
	Token          string
	ExpirationDate time.Time
}
