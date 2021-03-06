package service

import (
	"time"

	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/GG_Backend_tech_challenge/src/repository"
	"github.com/jinzhu/gorm"
)

type EventService interface {
	AddGuestToGuestList(name string, tableID int, accompany_guests int, TableRepo repository.TableRepository, GuestRepo repository.GuestRepository) (bool, model.RespondByName)
	GetGuests(GuestRepo repository.GuestRepository) (bool, model.RespondByGuests)
	AddGuestToArrivedGuests(name string, accompany_guests int, TableRepo repository.TableRepository, GuestRepo repository.GuestRepository, GuestArrivalsRepo repository.GuestArrivalsRepository) (bool, model.RespondByName)
	DeleteArrivedGuest(name string, TableRepo repository.TableRepository, GuestRepo repository.GuestRepository, GuestArrivalsRepo repository.GuestArrivalsRepository) bool
	GetArrivedGuests(GuestArrivalsRepo repository.GuestArrivalsRepository) (bool, model.RespondByArrivedGuests)
	GetEmptySeats(TableRepo repository.TableRepository) model.RespondBySeats
	AddTable(tableID int, capacity int, TableRepo repository.TableRepository) bool
}

type Eventservice struct {
	DB *gorm.DB
}

func (e *Eventservice) AddGuestToGuestList(name string, tableID int, accompany_guests int, TableRepo repository.TableRepository, GuestRepo repository.GuestRepository) (bool, model.RespondByName) {
	responseModel := model.RespondByName{Name: name}
	ok_guest, guest := GuestRepo.GetGuestByName(name)
	ok_table, table := TableRepo.DoesTableExist(tableID)

	if ok_guest || !ok_table {
		return false, responseModel
	}

	guest = model.Guest{Name: name, AccompanyingGuests: accompany_guests}

	ok_table, _ = TableRepo.AssignTableToGuest(table, &guest)

	if !ok_table {
		return false, responseModel
	}

	GuestRepo.AddGuest(guest)

	return true, responseModel

}

func (e *Eventservice) GetGuests(GuestRepo repository.GuestRepository) (bool, model.RespondByGuests) {
	ok, guests := GuestRepo.GetGuests()
	responseModel := model.RespondByGuests{
		Guests: guests,
	}

	if ok {
		return true, responseModel
	} else {
		return false, responseModel
	}
}

func (e *Eventservice) AddGuestToArrivedGuests(name string, accompany_guests int, TableRepo repository.TableRepository, GuestRepo repository.GuestRepository, GuestArrivalsRepo repository.GuestArrivalsRepository) (bool, model.RespondByName) {
	responseModel := model.RespondByName{Name: name}
	ok_guest, guest := GuestRepo.GetGuestByName(name)
	if !ok_guest {
		return false, responseModel
	}

	if guest.AccompanyingGuests < accompany_guests {
		increase := (accompany_guests + 1) - (guest.AccompanyingGuests + 1)
		ok_table, table := TableRepo.DoesTableExist(guest.TableID)

		if !ok_table {
			return false, responseModel
		}

		ok_table, _ = TableRepo.IncreaseGuestSeats(table, increase)

		if !ok_table {
			return false, responseModel
		}

		guest.AccompanyingGuests += increase

		ok_arrive := GuestArrivalsRepo.AddArrivedGuest(guest, time.Now())

		return ok_arrive, responseModel

	} else if guest.AccompanyingGuests > accompany_guests {
		decrease := (guest.AccompanyingGuests + 1) - (accompany_guests + 1)
		ok_table, table := TableRepo.DoesTableExist(guest.TableID)

		if !ok_table {
			return false, responseModel
		}

		ok_table, _ = TableRepo.DecreaseGuestSeats(table, decrease)

		if !ok_table {
			return false, responseModel
		}

		guest.AccompanyingGuests -= decrease

		ok_arrive := GuestArrivalsRepo.AddArrivedGuest(guest, time.Now())

		return ok_arrive, responseModel
	}

	ok_table, _ := TableRepo.DoesTableExist(guest.TableID)

	if !ok_table {
		return false, responseModel
	}

	ok_arrive := GuestArrivalsRepo.AddArrivedGuest(guest, time.Now())

	return ok_arrive, responseModel

}

func (e *Eventservice) DeleteArrivedGuest(name string, TableRepo repository.TableRepository, GuestRepo repository.GuestRepository, GuestArrivalsRepo repository.GuestArrivalsRepository) bool {
	ok_guest, guest := GuestRepo.GetGuestByName(name)
	if !ok_guest {
		return false
	}

	ok_table, table := TableRepo.DoesTableExist(guest.TableID)

	if !ok_table {
		return false
	}

	ok_table, _ = TableRepo.RemoveGuestFromTable(guest, table)

	if !ok_table {
		return false
	}

	ok_guest = GuestArrivalsRepo.DeleteArrivedGuestByGuestName(guest.Name)

	return ok_guest

}

func (e *Eventservice) GetArrivedGuests(GuestArrivalsRepo repository.GuestArrivalsRepository) (bool, model.RespondByArrivedGuests) {
	ok, arrivedguests := GuestArrivalsRepo.GetArrivedGuests()

	responseModel := model.RespondByArrivedGuests{Guests: arrivedguests}

	return ok, responseModel
}

func (e *Eventservice) GetEmptySeats(TableRepo repository.TableRepository) model.RespondBySeats {
	seats := TableRepo.GetEmptySeats()
	reponseModel := model.RespondBySeats{
		SeatsEmpty: seats,
	}
	return reponseModel
}

func (e *Eventservice) AddTable(tableID int, capacity int, TableRepo repository.TableRepository) bool {
	ok, _ := TableRepo.DoesTableExist(tableID)
	if ok {
		return false
	} else {
		return TableRepo.AddTable(tableID, capacity)
	}
}
