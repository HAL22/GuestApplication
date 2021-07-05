package model

import (
	"time"
)

type GuestArrivals struct {
	ID                 int       `gorm:"AUTO_INCREMENT" gorm:"column:id" json:"-"`
	Name               string    `json:"name"`
	AccompanyingGuests int       `json:"accompanying_guests"`
	ArrivalTime        time.Time `json:"time_arrived"`
	GuestID            int       `json:"-"`
}
