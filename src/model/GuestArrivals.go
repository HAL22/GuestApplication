package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type GuestArrivals struct {
	gorm.Model
	Name               string
	AccompanyingGuests int
	ArrivalTime        time.Time
	GuestID            int
}
