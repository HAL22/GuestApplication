package model

type RespondByName struct {
	Name string `json:"name"`
}

type RespondByGuests struct {
	Guests []Guest
}

type RespondByArrivedGuests struct {
	Guests []GuestArrivals
}

type RespondBySeats struct {
	SeatsEmpty int `json:"seats_empty"`
}

type ResponseAddGuest struct {
	Table            int `json:"table"`
	Accompany_guests int `json:"accompany_guests"`
}

type ResponseAddArrivedGuest struct {
	Accompany_guests int `json:"accompany_guests"`
}

type ResponseTable struct {
	Table    int `json:"table"`
	Capacity int `json:"capacity"`
}
