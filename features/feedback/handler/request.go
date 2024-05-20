package handler

type FeedbackRequest struct {
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Rating     int    `json:"rating" form:"rating"`
	Feedback   string `json:"feedback" form:"feedback"`
}
