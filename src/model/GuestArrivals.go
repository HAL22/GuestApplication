package model

import (
	"time"

	"gorm.io/gorm"
)

type GuestArrivals struct {
	gorm.Model
	ID           uint
	ArrivedGuest Guest
	GuestID      uint
	ArrivalTime  time.Time
}
