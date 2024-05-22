package data

import (
	"errors"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation"
	"gorm.io/gorm"
)

type reservationQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservation.DataInterface {
	return &reservationQuery{
		db: db,
	}
}

func (r *reservationQuery) CheckAvailability(input reservation.Core) error {

	var reservation []Reservation

	tx := r.db.Where("homestay_id = ?", input.HomestayID).
		Where("(? >= start_date AND ? <= end_date) OR (? >= start_date AND ? <= end_date)", input.StartDate, input.StartDate, input.EndDate, input.EndDate).
		Find(&reservation)
	if tx.Error != nil {
		return tx.Error
	}
	if len(reservation) == 0 {
		return nil
	}

	return errors.New("not available")
}

func (r *reservationQuery) CreateReservation(input reservation.Core) error {
	var reservationGorm Reservation

	reservationGorm = Reservation{
		Model:      gorm.Model{},
		UserID:     input.UserID,
		HomestayID: input.HomestayID,
		StartDate:  input.StartDate,
		EndDate:    input.EndDate,
		TotalPrice: input.TotalPrice,
		Status:     input.Status,
	}

	tx := r.db.Create(&reservationGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *reservationQuery) GetHistory(UserId uint) (data []reservation.Core, err error) {
	var historyReservationGorm []Reservation
	tx := r.db.Where("user_id = ?", UserId).Find(&historyReservationGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var historyReservationCore []reservation.Core
	for _, v := range historyReservationGorm {
		historyReservationCore = append(historyReservationCore, reservation.Core{
			ID:         v.ID,
			UserID:     v.UserID,
			HomestayID: v.HomestayID,
			StartDate:  v.StartDate,
			EndDate:    v.EndDate,
			TotalPrice: v.TotalPrice,
			Status:     v.Status,
		})
	}

	return historyReservationCore, nil
}
