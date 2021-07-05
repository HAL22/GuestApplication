package repository

import (
	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/jinzhu/gorm"
)

type GuestRepository interface {
	GetGuestByName(name string) (bool, model.Guest)
	AddGuest(guest model.Guest) bool
	GetGuests() (bool, []model.Guest)
}

type GuestRepo struct {
	DB *gorm.DB
}

func (g *GuestRepo) GetGuestByName(name string) (bool, model.Guest) {
	var guest model.Guest
	result := g.DB.Where("name = ?", name).First(&guest)
	if result.Error == nil {
		return true, guest
	} else {
		return false, guest
	}
}

func (g *GuestRepo) AddGuest(guest model.Guest) bool {
	result := g.DB.Create(&guest)
	if result.Error == nil {
		return true
	} else {
		return false
	}
}

func (g *GuestRepo) GetGuests() (bool, []model.Guest) {
	var guests []model.Guest
	result := g.DB.Find(&guests)
	if result.Error == nil {
		return true, guests
	} else {
		return false, guests
	}
}

func NewGuestRepo(db *gorm.DB) GuestRepository {
	return &GuestRepo{DB: db}
}
