package service

import (
	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user"
)

type reservationService struct {
	reservationData reservation.DataInterface
	userData        user.DataInterface
	homestayData    homestay.DataInterface
}

func New(rd reservation.DataInterface, ud user.DataInterface, hd homestay.DataInterface) reservation.ServiceInterface {
	return &reservationService{
		reservationData: rd,
		userData:        ud,
		homestayData:    hd,
	}
}

func (r *reservationService) CheckAvailability(input reservation.Core) error {
	err := r.reservationData.CheckAvailability(input)
	if err != nil {
		return err
	}
	return nil
}

func (r *reservationService) GetHistory(UserId uint) (data []reservation.Core, err error) {
	return r.reservationData.GetHistory(UserId)
}

func (r *reservationService) CreateReservation(input reservation.Core) error {
	// check userid exist
	_, err := r.userData.SelectProfileById(input.UserID)
	if err != nil {
		return err
	}

	// check homestay exist
	homestay, err := r.homestayData.GetHomestayById(input.HomestayID)
	if err != nil {
		return err
	}

	// check homestay availability
	err = r.reservationData.CheckAvailability(input)
	if err != nil {
		return err
	}

	// calculate homestay price * price per night from homestay
	input.TotalPrice = int(input.EndDate.Sub(input.StartDate).Hours()/24) * int(homestay.CostPerNight)

	// set status
	input.Status = "success"

	return r.reservationData.CreateReservation(input)
}
