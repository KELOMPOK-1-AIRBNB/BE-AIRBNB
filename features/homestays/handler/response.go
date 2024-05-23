package handler

type HomestayResponse struct {
	ID            uint
	HomestayName  string
	Address       string
	PricePerNight int
	Description   string
	Images1       string
	Images2       string
	Images3       string
}

type HomestayResponseById struct {
	ID            uint
	HomestayName  string
	Address       string
	Description   string
	PricePerNight int
	Images1       string
	Images2       string
	Images3       string
}
