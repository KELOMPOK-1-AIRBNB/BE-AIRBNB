package handler

type HomestayRequest struct {
	HomestayName  string `json:"homestay_name" form:"homestay_name"`
	Description   string `json:"description" form:"description"`
	Address       string `json:"address" form:"address"`
	Images1       string `json:"images1" form:"images1"`
	Images2       string `json:"images2" form:"images2"`
	Images3       string `json:"images3" form:"images3"`
	PricePerNight int    `json:"price_per_night" form:"price_per_night"`
}
