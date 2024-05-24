package service

import (
	"errors"
	"testing"
	"time"

	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/reservation"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckAvailability(t *testing.T) {
	t.Run("success check availability", func(t *testing.T) {
		repoReservationMock := new(mocks.ReservationData)
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		startDate := time.Date(2024, time.May, 24, 0, 0, 0, 0, time.UTC)
		endDate := startDate.Add(48 * time.Hour)
		input := reservation.Core{
			UserID:     1,
			HomestayID: 1,
			StartDate:  startDate,
			EndDate:    endDate,
		}

		repoReservationMock.On("CheckAvailability", mock.Anything).Return(nil)

		srv := New(repoReservationMock, repoUserMock, repoHomestayMock)
		err := srv.CheckAvailability(input)

		assert.NoError(t, err)
		repoReservationMock.AssertExpectations(t)
	})

	t.Run("check availability failed due to invalid input", func(t *testing.T) {
		repoReservationMock := new(mocks.ReservationData)
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		startDate := time.Date(2024, time.May, 24, 0, 0, 0, 0, time.UTC)
		endDate := startDate.Add(48 * time.Hour)
		invalidInput := reservation.Core{
			UserID:     0,
			HomestayID: 1,
			StartDate:  startDate,
			EndDate:    endDate,
		}

		srv := New(repoReservationMock, repoUserMock, repoHomestayMock)
		err := srv.CheckAvailability(invalidInput)

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid input")
		repoReservationMock.AssertExpectations(t)
	})

	t.Run("check availability failed due to internal error", func(t *testing.T) {
		repoReservationMock := new(mocks.ReservationData)
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		startDate := time.Date(2024, time.May, 24, 0, 0, 0, 0, time.UTC)
		endDate := startDate.Add(48 * time.Hour)
		input := reservation.Core{
			UserID:     1,
			HomestayID: 1,
			StartDate:  startDate,
			EndDate:    endDate,
		}

		repoReservationMock.On("CheckAvailability", mock.Anything).Return(errors.New("internal error"))

		srv := New(repoReservationMock, repoUserMock, repoHomestayMock)
		err := srv.CheckAvailability(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "internal error")
		repoReservationMock.AssertExpectations(t)
	})
}

func TestCreateReservation(t *testing.T) {
	t.Run("success create reservation", func(t *testing.T) {
		repoReservationMock := new(mocks.ReservationData)
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		startDate := time.Date(2024, time.May, 24, 0, 0, 0, 0, time.UTC)
		endDate := startDate.Add(48 * time.Hour)
		input := reservation.Core{
			UserID:     1,
			HomestayID: 1,
			StartDate:  startDate,
			EndDate:    endDate,
		}

		repoUserMock.On("SelectProfileById", input.UserID).Return(&user.Core{}, nil)
		repoHomestayMock.On("GetHomestayById", input.HomestayID).Return(homestay.Core{}, nil)
		repoReservationMock.On("CheckAvailability", mock.Anything).Return(nil)
		repoReservationMock.On("CreateReservation", mock.Anything).Return(nil)

		srv := New(repoReservationMock, repoUserMock, repoHomestayMock)
		err := srv.CreateReservation(input)

		assert.NoError(t, err)
		repoUserMock.AssertExpectations(t)
		repoHomestayMock.AssertExpectations(t)
		repoReservationMock.AssertExpectations(t)
	})

	t.Run("create reservation failed due to invalid input", func(t *testing.T) {
		repoReservationMock := new(mocks.ReservationData)
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		startDate := time.Date(2024, time.May, 24, 0, 0, 0, 0, time.UTC)
		endDate := startDate.Add(48 * time.Hour)
		invalidInput := reservation.Core{
			UserID:     0,
			HomestayID: 1,
			StartDate:  startDate,
			EndDate:    endDate,
		}

		srv := New(repoReservationMock, repoUserMock, repoHomestayMock)
		err := srv.CreateReservation(invalidInput)

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid input")
		repoReservationMock.AssertExpectations(t)
	})

	t.Run("create reservation failed due to user not found", func(t *testing.T) {
		repoReservationMock := new(mocks.ReservationData)
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		startDate := time.Date(2024, time.May, 24, 0, 0, 0, 0, time.UTC)
		endDate := startDate.Add(48 * time.Hour)
		input := reservation.Core{
			UserID:     1,
			HomestayID: 1,
			StartDate:  startDate,
			EndDate:    endDate,
		}

		repoUserMock.On("SelectProfileById", input.UserID).Return(nil, errors.New("user not found"))

		srv := New(repoReservationMock, repoUserMock, repoHomestayMock)
		err := srv.CreateReservation(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "user not found")
		repoUserMock.AssertExpectations(t)
	})

	t.Run("create reservation failed due to homestay not found", func(t *testing.T) {
		repoReservationMock := new(mocks.ReservationData)
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		startDate := time.Date(2024, time.May, 24, 0, 0, 0, 0, time.UTC)
		endDate := startDate.Add(48 * time.Hour)
		input := reservation.Core{
			UserID:     1,
			HomestayID: 1,
			StartDate:  startDate,
			EndDate:    endDate,
		}
		repoUserMock.On("SelectProfileById", input.UserID).Return(&user.Core{}, nil)
		repoHomestayMock.On("GetHomestayById", input.HomestayID).Return(homestay.Core{}, errors.New("homestay not found"))

		srv := New(repoReservationMock, repoUserMock, repoHomestayMock)
		err := srv.CreateReservation(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "homestay not found")
		repoUserMock.AssertExpectations(t)
		repoHomestayMock.AssertExpectations(t)
	})

	t.Run("create reservation failed due to homestay is unavailable", func(t *testing.T) {
		repoReservationMock := new(mocks.ReservationData)
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		startDate := time.Date(2024, time.May, 24, 0, 0, 0, 0, time.UTC)
		endDate := startDate.Add(48 * time.Hour)
		input := reservation.Core{
			UserID:     1,
			HomestayID: 1,
			StartDate:  startDate,
			EndDate:    endDate,
		}
		repoUserMock.On("SelectProfileById", input.UserID).Return(&user.Core{}, nil)
		repoHomestayMock.On("GetHomestayById", input.HomestayID).Return(homestay.Core{}, nil)
		repoReservationMock.On("CheckAvailability", mock.Anything).Return(errors.New("homestay unavailable"))

		srv := New(repoReservationMock, repoUserMock, repoHomestayMock)
		err := srv.CreateReservation(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "homestay unavailable")
		repoUserMock.AssertExpectations(t)
		repoHomestayMock.AssertExpectations(t)
		repoReservationMock.AssertExpectations(t)
	})
}

func TestGetHistory(t *testing.T) {
	t.Run("success get reservation history", func(t *testing.T) {
		repoReservationData := new(mocks.ReservationData)

		expectedData := []reservation.Core{
			{ID: 1, UserID: 1, HomestayID: 1},
			{ID: 1, UserID: 1, HomestayID: 2},
		}

		repoReservationData.On("GetHistory", uint(1)).Return(expectedData, nil)

		srv := New(repoReservationData, nil, nil)
		result, err := srv.GetHistory(uint(1))

		assert.NoError(t, err)
		assert.Equal(t, expectedData, result)
		repoReservationData.AssertExpectations(t)
	})

	t.Run("fail to get reservation history", func(t *testing.T) {
		repoReservationData := new(mocks.ReservationData)

		repoReservationData.On("GetHistory", uint(1)).Return(nil, errors.New("database error"))

		srv := New(repoReservationData, nil, nil)
		result, err := srv.GetHistory(uint(1))

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
		assert.Nil(t, result)
		repoReservationData.AssertExpectations(t)
	})
}
