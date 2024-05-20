package handler

type HomestayResponse struct {
	ID            uint
	HomestayName  string
	Address       string
	PricePerNight int
}

type HomestayResponseById struct {
	ID           uint
	HomestayName string
	Description  string
}
