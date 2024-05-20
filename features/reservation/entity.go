package reservation

import (
	"time"
)

type Core struct {
	ID         uint
	UserID     uint
	HomestayID uint
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DataInterface interface {
	CheckAvailability(input Core) error
	CreateReservation(input Core) error
	GetHistory(UserId uint) (data []Core, err error)
}

type ServiceInterface interface {
	CheckAvailability(input Core) error
	GetHistory(UserId uint) (data []Core, err error)
	CreateReservation(input Core) error
}
