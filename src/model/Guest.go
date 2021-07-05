package model

type Guest struct {
	ID                 int    `gorm:"AUTO_INCREMENT" gorm:"column:id" json:"-"`
	Name               string `json:"name"`
	AccompanyingGuests int    `json:"accompanying_guests"`
	TableID            int    `json:"table"`
}
