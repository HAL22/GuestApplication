package model

type Table struct {
	ID           uint
	Guests       []Guest `gorm:"foreignKey:TableID"  json:"guest"`
	Capacity     int
	Sizeofguests int
	Emptyseats   int
}
