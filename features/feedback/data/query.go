package data

import (
	"errors"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback"
	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.DataInterface {
	return &feedbackQuery{
		db: db,
	}
}

func (f *feedbackQuery) CreateFeedback(input feedback.Core) error {
	var count int64
	tx := f.db.Model(&Feedback{}).Where("user_id = ? AND homestay_id = ?", input.UserID, input.HomestayID).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}

	if count > 0 {
		return errors.New("feedback already exist")
	}

	var feedbackGorm Feedback

	feedbackGorm = Feedback{
		Model:      gorm.Model{},
		UserID:     input.UserID,
		HomestayID: input.HomestayID,
		Rating:     input.Rating,
		Feedback:   input.Feedback,
	}

	tx = f.db.Create(&feedbackGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (f *feedbackQuery) GetFeedbackByHomestayId(homestayId uint) (data []feedback.Core, err error) {
	var feedbackGorm []Feedback
	tx := f.db.Where("homestay_id = ?", homestayId).Find(&feedbackGorm)
	if tx.Error != nil {
		return data, tx.Error
	}

	var feedbackCore []feedback.Core
	for _, v := range feedbackGorm {
		feedbackCore = append(feedbackCore, feedback.Core{
			ID:         v.ID,
			UserID:     v.UserID,
			HomestayID: v.HomestayID,
			Rating:     v.Rating,
			Feedback:   v.Feedback,
			CreatedAt:  v.CreatedAt,
		})
	}

	return feedbackCore, nil
}
