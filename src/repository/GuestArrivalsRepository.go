package repository

import (
	"time"

	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/jinzhu/gorm"
)

type GuestArrivalsRepository interface {
	AddArrivedGuest(guest model.Guest, time time.Time) bool
	DeleteArrivedGuestByGuestName(name string) bool
	GetArrivedGuestByGuestName(name string) (bool, model.GuestArrivals)
	GetArrivedGuests() (bool, []model.GuestArrivals)
}

type GuestArrivalsRepo struct {
	DB *gorm.DB
}

func (g *GuestArrivalsRepo) AddArrivedGuest(guest model.Guest, time time.Time) bool {
	arrivedguest := model.GuestArrivals{
		Name:               guest.Name,
		AccompanyingGuests: guest.AccompanyingGuests,
		ArrivalTime:        time,
		GuestID:            guest.ID,
	}
	result := g.DB.Create(&arrivedguest)
	if result.Error == nil {
		return true
	} else {
		return false
	}
}

func (g *GuestArrivalsRepo) DeleteArrivedGuestByGuestName(name string) bool {
	result := g.DB.Unscoped().Where("name = ?", name).Delete(&model.GuestArrivals{})
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}

func (g *GuestArrivalsRepo) GetArrivedGuestByGuestName(name string) (bool, model.GuestArrivals) {
	var arrivedguest model.GuestArrivals
	result := g.DB.Where("name = ?", name).First(&arrivedguest)
	if result.Error == nil {
		return true, arrivedguest
	} else {
		return false, arrivedguest
	}
}

func (g *GuestArrivalsRepo) GetArrivedGuests() (bool, []model.GuestArrivals) {
	var arrivedguests []model.GuestArrivals
	result := g.DB.Find(&arrivedguests)
	if result.Error == nil {
		return true, arrivedguests
	} else {
		return false, arrivedguests
	}
}

func NewGuestArrival(db *gorm.DB) GuestArrivalsRepository {
	return &GuestArrivalsRepo{db}
}
