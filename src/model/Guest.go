package model

type Guest struct {
	ID                 int `gorm:"AUTO_INCREMENT" gorm:"column:id" json:"id"`
	Name               string
	AccompanyingGuests int
	TableID            int
}
