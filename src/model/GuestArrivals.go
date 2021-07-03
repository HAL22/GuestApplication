package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type GuestArrivals struct {
	gorm.Model

	ArrivedGuest Guest `gorm:"foreignKey:GuestArrivalID"`
	ArrivalTime  time.Time
}
