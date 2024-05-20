package handler

import "time"

type AvailableResponse struct {
	Status string `json:"status"`
}

type HistoryResponse struct {
	UserID     uint      `json:"user_id"`
	HomestayID uint      `json:"homestay_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	TotalPrice int       `json:"total_price"`
	Status     string    `json:"status"`
}
