package data

import (
	userData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user/data"

	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	UserID        uint
	HomestayName  string
	Address       string
	Images1       string
	Images2       string
	Images3       string
	Description   string
	PricePerNight int
	User          userData.User `gorm:"foreignKey:UserID"`
}
