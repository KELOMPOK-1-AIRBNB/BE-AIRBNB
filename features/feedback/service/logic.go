package service

import (
	"errors"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback"
	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation"
)

type feedbackService struct {
	feedbackData    feedback.DataInterface
	reservationData reservation.DataInterface
	homestay        homestay.DataInterface
}

func New(fd feedback.DataInterface, rd reservation.DataInterface, homeStay homestay.DataInterface) feedback.ServiceInterface {
	return &feedbackService{
		feedbackData:    fd,
		reservationData: rd,
		homestay:        homeStay,
	}
}

func (f *feedbackService) CreateFeedback(input feedback.Core) error {
	if input.UserID <= 0 || input.HomestayID <= 0 || input.Feedback == "" || input.Rating <= 0 || input.Rating > 5 {
		return errors.New("invalid input")
	}

	err := f.feedbackData.CreateFeedback(input)
	if err != nil {
		return err
	}
	return nil
}

func (f *feedbackService) GetFeedbackByHomestayId(homestayId uint) (data []feedback.Core, err error) {
	_, err = f.homestay.GetHomestayById(homestayId)
	if err != nil {
		return data, errors.New("homestay not found")
	}

	return f.feedbackData.GetFeedbackByHomestayId(homestayId)
}
