package model

type RespondByName struct {
	Name string
}

type RespondByGuests struct {
	Guests []Guest
}

type RespondByArrivedGuests struct {
	Guests []GuestArrivals
}

type RespondBySeats struct {
	SeatsEmpty int
}

type ResponseAddGuest struct {
	Table            int
	Accompany_guests int
}
