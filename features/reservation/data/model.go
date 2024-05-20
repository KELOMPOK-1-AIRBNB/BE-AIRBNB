package data

import (
	homestayData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/data"
	userData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user/data"
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID     uint
	HomestayID uint
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice int
	Status     string
	User       userData.User         `gorm:"foreignKey:UserID"`
	Homestay   homestayData.Homestay `gorm:"foreignKey:HomestayID"`
}
