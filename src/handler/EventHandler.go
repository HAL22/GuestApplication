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
	guestRepo := repository.GuestRepo{
		DB: e.DB,
	}
	tableRepo := repository.TableRepo{
		DB: e.DB,
	}
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
