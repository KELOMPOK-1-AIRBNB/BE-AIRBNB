package data

import (
	homestayData "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays/data"
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
	var homestay []homestayData.Homestay
	tx := r.db.Not("booked_start BETWEEN ? AND ? AND booked_end BETWEEN ? AND ?", input.StartDate, input.EndDate, input.StartDate, input.EndDate).Where("homestay_id = ?", input.HomestayID).Find(&homestay)
	if tx.Error != nil {
		return tx.Error
	}
	if len(homestay) > 0 {
		return tx.Error
	}
	return nil
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

	updateHomestay := Reservation{
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	}

	tx2 := r.db.Model(&homestayData.Homestay{}).Where("id = ?", input.HomestayID).Updates(updateHomestay)
	if tx2.Error != nil {
		return tx2.Error
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
