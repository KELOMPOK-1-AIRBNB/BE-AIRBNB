package homestay

import "time"

type TaskListResponseCore struct {
	ID              uint
	TaskName        string
	DescriptionTask string
	StatusTask      string
}

type Core struct {
	ID           uint
	UserID       uint
	HomestayName string
	Description  string
	Address      string
	Images1      string
	Images2      string
	Images3      string
	StartDate    time.Time
	EndDate      time.Time
	CostPerNight int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAll(id uint) ([]Core, error)
	SelectAllForUser() ([]Core, error)
	GetHomestayById(id uint) (Core, error)
	Delete(id uint) error
	Update(id uint, input Core) error
	GetUserByHomestayId(id uint) (Core, error)
	GetMyHomestay(id uint) ([]Core, error)
}

type ServiceInterface interface {
	Create(input Core) error
	GetAll(id uint) ([]Core, error)
	GetAllForUser(id uint) ([]Core, error)
	GetHomestayById(id uint, idUser uint) (Core, error)
	Delete(id uint, idUser uint) error
	Update(id uint, idUser uint, input Core) error
	GetMyHomestay(id uint) ([]Core, error)
}
