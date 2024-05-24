package service

import (
	"errors"
	"testing"

	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/feedback"
	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateFeedback(t *testing.T) {
	t.Run("success create feedback", func(t *testing.T) {
		repoFeedbackMock := new(mocks.FeedbackData)
		input := feedback.Core{
			UserID:     1,
			HomestayID: 1,
			Rating:     5,
			Feedback:   "cozy room",
		}

		repoFeedbackMock.On("CreateFeedback", mock.Anything).Return(nil)
		srv := New(repoFeedbackMock, nil, nil)
		err := srv.CreateFeedback(input)

		assert.NoError(t, err)
		repoFeedbackMock.AssertExpectations(t)
	})

	t.Run("fail create feedback due to invalid input", func(t *testing.T) {
		repoFeedbackMock := new(mocks.FeedbackData)

		invalidInput := feedback.Core{
			UserID:     1,
			HomestayID: 0,
			Rating:     5,
			Feedback:   "cozy room",
		}

		srv := New(repoFeedbackMock, nil, nil)
		err := srv.CreateFeedback(invalidInput)

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid input")
	})

	t.Run("fail create feedback due to server error", func(t *testing.T) {
		repoFeedbackMock := new(mocks.FeedbackData)

		input := feedback.Core{
			UserID:     1,
			HomestayID: 1,
			Rating:     5,
			Feedback:   "cozy room",
		}

		repoFeedbackMock.On("CreateFeedback", mock.Anything).Return(errors.New("internal server error"))
		srv := New(repoFeedbackMock, nil, nil)
		err := srv.CreateFeedback(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "internal server error")
		repoFeedbackMock.AssertExpectations(t)
	})
}

func TestGetFeedbackByHomestayId(t *testing.T) {
	t.Run("success get feedback by homestay id", func(t *testing.T) {
		repoFeedbackMock := new(mocks.FeedbackData)
		repoHomestayMock := new(mocks.HomestayData)

		expectedFeedbacks := []feedback.Core{
			{ID: 1, HomestayID: uint(1), UserID: 1, Feedback: "Great place!", Rating: 5},
			{ID: 2, HomestayID: uint(1), UserID: 2, Feedback: "Nice stay", Rating: 4},
		}
		repoHomestayMock.On("GetHomestayById", uint(1)).Return(homestay.Core{ID: uint(1)}, nil)
		repoFeedbackMock.On("GetFeedbackByHomestayId", uint(1)).Return(expectedFeedbacks, nil)

		srv := New(repoFeedbackMock, nil, repoHomestayMock)
		result, err := srv.GetFeedbackByHomestayId(uint(1))

		assert.NoError(t, err)
		assert.Equal(t, expectedFeedbacks, result)
		repoHomestayMock.AssertExpectations(t)
		repoFeedbackMock.AssertExpectations(t)
	})

	t.Run("fail get feedback by homestay id - homestay not found", func(t *testing.T) {
		repoFeedbackMock := new(mocks.FeedbackData)
		repoHomestayMock := new(mocks.HomestayData)

		repoHomestayMock.On("GetHomestayById", uint(1)).Return(homestay.Core{}, errors.New("homestay not found"))

		srv := New(repoFeedbackMock, nil, repoHomestayMock)
		result, err := srv.GetFeedbackByHomestayId(uint(1))

		assert.Error(t, err)
		assert.EqualError(t, err, "homestay not found")
		assert.Nil(t, result)
		repoHomestayMock.AssertExpectations(t)
		repoFeedbackMock.AssertExpectations(t)
	})
}
