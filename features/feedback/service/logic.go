package service

import (
	"errors"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback"
	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user"
)

type feedbackService struct {
	feedbackData feedback.DataInterface
	userData     user.DataInterface
	homeStay     homestay.DataInterface
}

func New(fd feedback.DataInterface, ud user.DataInterface, hd homestay.DataInterface) feedback.ServiceInterface {
	return &feedbackService{
		feedbackData: fd,
		userData:     ud,
		homeStay:     hd,
	}
}

func (f *feedbackService) CreateFeedback(input feedback.Core) error {
	if input.UserID <= 0 || input.HomestayID <= 0 || input.Feedback == "" || input.Rating <= 0 || input.Rating > 5 {
		return errors.New("invalid input")
	}

	_, err := f.userData.SelectProfileById(input.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	_, err = f.homeStay.GetHomestayById(input.HomestayID)
	if err != nil {
		return errors.New("homestay not found")
	}

	err = f.feedbackData.CreateFeedback(input)
	if err != nil {
		return err
	}
	return nil
}

func (f *feedbackService) GetFeedbackByHomestayId(homestayId uint) (data []feedback.Core, err error) {
	_, err = f.homeStay.GetHomestayById(homestayId)
	if err != nil {
		return data, errors.New("homestay not found")
	}

	return f.feedbackData.GetFeedbackByHomestayId(homestayId)
}
