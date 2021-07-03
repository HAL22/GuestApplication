package repository

import (
	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/jinzhu/gorm"
)

type TableRepository interface {
	AddTable(capacity int, tableID int) bool
	DoesTableExist(tableID int) (bool, model.Table)
	AssignTableToGuest(table model.Table, guest *model.Guest) (bool, int)
	RemoveGuestFromTable(table model.Table, guestName string) (bool, int)
	IncreaseGuestSeats(table model.Table, increase int) (bool, int)
	GetEmptySeats() int
}

type tableRepository struct {
	DB *gorm.DB
}

func (t tableRepository) AddTable(capacity int, tableID uint) bool {
	table := model.Table{
		ID:       tableID,
		Guests:   make([]model.Guest, 0, capacity),
		Capacity: capacity, Sizeofguests: 0,
		Emptyseats: capacity}

	result := t.DB.Create(&table)

	if result.Error == nil {
		return true
	} else {
		return false
	}

}

func (t tableRepository) DoesTableExist(tableID int) (bool, model.Table) {
	var table model.Table
	result := t.DB.Where("id = ?", tableID).First(&table)

	if result.Error == nil && table.ID == uint(tableID) {
		return true, table
	} else {
		return false, table
	}
}

func (t tableRepository) AssignTableToGuest(table model.Table, guest *model.Guest) (bool, int) {
	if guest.AccompanyingGuests+1 > table.Emptyseats {
		return false, 0
	} else {
		table.Emptyseats -= (guest.AccompanyingGuests + 1)
		table.Sizeofguests += (guest.AccompanyingGuests + 1)
		guest.TableID = int(table.ID)
		table.Guests = append(table.Guests, *guest)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("guests", table.Guests)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("sizeofguests", table.Sizeofguests)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("emptyseats", table.Emptyseats)
		return true, table.Emptyseats

	}

}

func (t tableRepository) RemoveGuestFromTable(table model.Table, guestName string) (bool, int) {
	var guest model.Guest
	var guestlist []model.Guest
	found := false
	for _, g := range table.Guests {
		if g.Name == guestName {
			found = true
			guest = g
			break
		}
	}
	if found {
		guestlist = make([]model.Guest, 0, table.Capacity-1)
		for _, g := range table.Guests {
			if g.Name != guestName {
				guestlist = append(guestlist, g)
			}
		}
		table.Emptyseats += (guest.AccompanyingGuests + 1)
		table.Sizeofguests -= (guest.AccompanyingGuests + 1)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("guests", guestlist)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("sizeofguests", table.Sizeofguests)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("emptyseats", table.Emptyseats)
		return true, table.Emptyseats

	} else {
		return false, table.Emptyseats
	}
}

func (t tableRepository) IncreaseGuestSeats(table model.Table, increase int) (bool, int) {
	if increase > table.Emptyseats {
		return false, 0
	} else {
		table.Emptyseats -= increase
		table.Sizeofguests += increase
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("sizeofguests", table.Sizeofguests)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("emptyseats", table.Emptyseats)
		return true, table.Emptyseats
	}

}

func (t tableRepository) GetEmptySeats() int {
	var sum int
	t.DB.Model(&model.Guest{}).Select("sum(emptyseats) as total").Find(&sum)
	return sum
}
