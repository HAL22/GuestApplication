package model

import (
	"github.com/jinzhu/gorm"
)

type Guest struct {
	gorm.Model
	Name               string
	AccompanyingGuests int
	TableID            int
	GuestArrivalID     int
}
