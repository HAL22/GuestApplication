package model

import (
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	ID           uint
	Guests       []Guest
	Capacity     int
	SizeOfGuests int
}
