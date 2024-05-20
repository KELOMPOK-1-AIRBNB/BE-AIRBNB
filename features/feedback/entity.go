package feedback

import (
	"time"
)

type Core struct {
	ID         uint
	UserID     uint
	HomestayID uint
	Rating     int
	Feedback   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DataInterface interface {
	CreateFeedback(input Core) error
	GetFeedbackByHomestayId(homestayId uint) (data []Core, err error)
}

type ServiceInterface interface {
	CreateFeedback(input Core) error
	GetFeedbackByHomestayId(homestayId uint) (data []Core, err error)
}
