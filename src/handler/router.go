package handler

import "github.com/gorilla/mux"

func NewRouter(eventHandler EventHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/guest_list/{name}", eventHandler.AddGuest).Methods("POST")

	r.HandleFunc("/guest_list", eventHandler.GetGuests).Methods("GET")

	r.HandleFunc("/guests/{name}", eventHandler.AddGuestToArrivedGuests).Methods("PUT")

	r.HandleFunc("/guests/{name}", eventHandler.DeleteArrivedGuest).Methods("DELETE")

	r.HandleFunc("/guests", eventHandler.GetArrivedGuests).Methods("GET")

	r.HandleFunc("/seats_empty", eventHandler.GetEmptySeats).Methods("GET")

	return r

}
