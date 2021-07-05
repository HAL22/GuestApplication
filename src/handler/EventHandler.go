package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"log"

	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/GG_Backend_tech_challenge/src/repository"
	"github.com/GG_Backend_tech_challenge/src/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type EventHandler struct {
	DB *gorm.DB
}

func (e *EventHandler) AddGuest(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())

	}

	var response model.ResponseAddGuest
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err.Error())

	}
	ServiceObject := service.Eventservice{
		DB: e.DB,
	}
	guestRepo := repository.NewGuestRepo(e.DB)
	tableRepo := repository.NewTableRepo(e.DB)
	log.Println(response)
	ok, result := ServiceObject.AddGuestToGuestList(name, response.Table, response.Accompany_guests, tableRepo, guestRepo)
	if ok {
		j, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(j)
		if err != nil {
			log.Panicln(err)
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
	}
}

func (e EventHandler) GetGuests(w http.ResponseWriter, req *http.Request) {
	ServiceObject := service.Eventservice{
		DB: e.DB,
	}
	guestRepo := repository.NewGuestRepo(e.DB)
	ok, result := ServiceObject.GetGuests(guestRepo)
	if ok {
		j, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(j)
		if err != nil {
			log.Panicln(err)
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
	}
}

func (e *EventHandler) AddGuestToArrivedGuests(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())

	}
	var response model.ResponseAddArrivedGuest
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err.Error())

	}
	ServiceObject := service.Eventservice{
		DB: e.DB,
	}
	guestRepo := repository.NewGuestRepo(e.DB)
	tableRepo := repository.NewTableRepo(e.DB)
	guestArrivalsRepo := repository.NewGuestArrival(e.DB)
	ok, result := ServiceObject.AddGuestToArrivedGuests(name, response.Accompany_guests, tableRepo, guestRepo, guestArrivalsRepo)
	if ok {
		j, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(j)
		if err != nil {
			log.Panicln(err)
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
	}
}

func (e *EventHandler) DeleteArrivedGuest(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	_, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())

	}
	ServiceObject := service.Eventservice{
		DB: e.DB,
	}
	guestRepo := repository.NewGuestRepo(e.DB)
	tableRepo := repository.NewTableRepo(e.DB)
	guestArrivalsRepo := repository.NewGuestArrival(e.DB)
	ok := ServiceObject.DeleteArrivedGuest(name, tableRepo, guestRepo, guestArrivalsRepo)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
	}

}

func (e *EventHandler) GetArrivedGuests(w http.ResponseWriter, req *http.Request) {
	ServiceObject := service.Eventservice{
		DB: e.DB,
	}
	guestArrivalsRepo := repository.NewGuestArrival(e.DB)
	ok, result := ServiceObject.GetArrivedGuests(guestArrivalsRepo)
	if ok {
		j, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(j)
		if err != nil {
			log.Panicln(err)
			return
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
	}

}

func (e *EventHandler) GetEmptySeats(w http.ResponseWriter, req *http.Request) {

	ServiceObject := service.Eventservice{
		DB: e.DB,
	}
	tableRepo := repository.NewTableRepo(e.DB)
	result := ServiceObject.GetEmptySeats(tableRepo)

	j, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(j)
	if err != nil {
		log.Panicln(err)
		return
	}
}

func (e *EventHandler) AddTable(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())

	}
	var response model.ResponseTable
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err.Error())

	}
	ServiceObject := service.Eventservice{
		DB: e.DB,
	}
	tableRepo := repository.NewTableRepo(e.DB)
	ok := ServiceObject.AddTable(response.Table, response.Capacity, tableRepo)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
	}

}
