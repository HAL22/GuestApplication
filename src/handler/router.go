package handler

import "github.com/gorilla/mux"

func NewRouter(eventHandler EventHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/guest_list/{name}", eventHandler.AddGuest).Methods("POST")

	return r

}
