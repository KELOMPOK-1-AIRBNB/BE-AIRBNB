package handler

import "time"

type FeedbackResponse struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	HomestayID uint      `json:"homestay_id"`
	Rating     int       `json:"rating"`
	Feedback   string    `json:"feedback"`
	CreatedAt  time.Time `json:"created_at"`
}
