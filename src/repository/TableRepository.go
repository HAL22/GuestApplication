package repository

import (
	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/jinzhu/gorm"
)

type TableRepository interface {
	AddTable(capacity int, tableID int) bool
	DoesTableExist(tableID int) (bool, model.Table)
	AssignTableToGuest(table model.Table, guest *model.Guest) (bool, int)
	RemoveGuestFromTable(guest model.Guest, table model.Table) (bool, int)
	IncreaseGuestSeats(table model.Table, increase int) (bool, int)
	DecreaseGuestSeats(table model.Table, decrease int) (bool, int)
	GetEmptySeats() int
}

type TableRepo struct {
	DB *gorm.DB
}

func (t *TableRepo) AddTable(tableID int, capacity int) bool {
	table := model.Table{
		ID:       uint(tableID),
		Capacity: capacity, Sizeofguests: 0,
		Emptyseats: capacity}

	result := t.DB.Create(&table)

	if result.Error == nil {
		return true
	} else {
		return false
	}

}

func (t *TableRepo) DoesTableExist(tableID int) (bool, model.Table) {
	var table model.Table
	result := t.DB.Where("id = ?", tableID).First(&table)

	if result.Error == nil && table.ID == uint(tableID) {
		return true, table
	} else {
		return false, table
	}
}

func (t *TableRepo) AssignTableToGuest(table model.Table, guest *model.Guest) (bool, int) {
	if guest.AccompanyingGuests+1 > table.Emptyseats {
		return false, 0
	} else {
		table.Emptyseats -= (guest.AccompanyingGuests + 1)
		table.Sizeofguests += (guest.AccompanyingGuests + 1)
		guest.TableID = int(table.ID)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("sizeofguests", table.Sizeofguests)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("emptyseats", table.Emptyseats)
		return true, table.Emptyseats

	}

}

func (t *TableRepo) RemoveGuestFromTable(guest model.Guest, table model.Table) (bool, int) {

	table.Emptyseats += (guest.AccompanyingGuests + 1)
	table.Sizeofguests -= (guest.AccompanyingGuests + 1)
	t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("sizeofguests", table.Sizeofguests)
	t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("emptyseats", table.Emptyseats)
	t.DB.Model(&model.Guest{}).Where("name = ?", guest.Name).Update("table_id", -1)
	return true, table.Emptyseats
}

func (t *TableRepo) IncreaseGuestSeats(table model.Table, increase int) (bool, int) {
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

func (t *TableRepo) DecreaseGuestSeats(table model.Table, decrease int) (bool, int) {
	if decrease <= 0 || table.Sizeofguests-decrease < 0 {
		return false, 0
	} else {
		table.Emptyseats += decrease
		table.Sizeofguests -= decrease
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("sizeofguests", table.Sizeofguests)
		t.DB.Model(&model.Table{}).Where("id = ?", table.ID).Update("emptyseats", table.Emptyseats)
		return true, table.Emptyseats
	}
}

func (t *TableRepo) GetEmptySeats() int {
	var sum int
	t.DB.Table("tables").Select("sum(emptyseats)").Row().Scan(&sum)
	return sum
}

func NewTableRepo(db *gorm.DB) TableRepository {
	return &TableRepo{db}
}
