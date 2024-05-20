package handler

import "time"

type AvailableRequest struct {
	HomestayID uint      `json:"homestay_id" form:"homestay_id"`
	StartDate  time.Time `json:"start_date" form:"start_date"`
	EndDate    time.Time `json:"end_date" form:"end_date"`
}

type ReservationRequest struct {
	HomestayID uint      `json:"homestay_id" form:"homestay_id"`
	StartDate  time.Time `json:"start_date" form:"start_date"`
	EndDate    time.Time `json:"end_date" form:"end_date"`
}
