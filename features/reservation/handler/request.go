package handler

type AvailableRequest struct {
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
}

type ReservationRequest struct {
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
}
