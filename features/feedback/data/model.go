package data

import (
	homestayData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/data"
	userData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user/data"
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	UserID     uint
	HomestayID uint
	Rating     int
	Feedback   string
	User       userData.User         `gorm:"foreignKey:UserID"`
	Homestay   homestayData.Homestay `gorm:"foreignKey:HomestayID"`
}
