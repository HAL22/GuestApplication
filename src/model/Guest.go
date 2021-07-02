package model

import (
	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	ID                 uint
	Name               string
	AccompanyingGuests int
	GuestTable         Table
	TableID            uint
}
