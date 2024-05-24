package service

import (
	"errors"
	"testing"

	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {

	t.Run("success create homestay", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)
		repoUserMock := new(mocks.UserData)
		input := homestay.Core{
			UserID:       1,
			HomestayName: "lotus",
			Description:  "1 queen size bed",
			Address:      "bandung",
			Images1:      "http//google.com/images1",
			Images2:      "http//google.com/images2",
			Images3:      "http//google.com/images3",
			CostPerNight: 150000,
		}

		repoUserMock.On("SelectProfileById", uint(1)).Return(&user.Core{Role: "host"}, nil)
		repoHomestayMock.On("Insert", mock.Anything).Return(nil)
		srv := New(repoHomestayMock, repoUserMock)

		err := srv.Create(input)
		assert.NoError(t, err)
		repoHomestayMock.AssertExpectations(t)
		repoUserMock.AssertExpectations(t)
	})

	t.Run("failed create homestay due to not host", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)
		repoUserMock := new(mocks.UserData)

		input := homestay.Core{
			UserID:       1,
			HomestayName: "lotus",
			Description:  "1 queen size bed",
			Address:      "bandung",
			Images1:      "http//google.com/images1",
			Images2:      "http//google.com/images2",
			Images3:      "http//google.com/images3",
			CostPerNight: 150000,
		}

		repoUserMock.On("SelectProfileById", uint(1)).Return(&user.Core{Role: "user"}, nil)

		srv := New(repoHomestayMock, repoUserMock)
		err := srv.Create(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "you're not host. make a host first")
		repoUserMock.AssertExpectations(t)
	})
	t.Run("failed create homestay due to invalid input", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)
		repoUserMock := new(mocks.UserData)
		invalidInput := homestay.Core{}

		srv := New(repoHomestayMock, repoUserMock)
		err := srv.Create(invalidInput)

		assert.Error(t, err)
		assert.EqualError(t, err, "all list must be filled")
	})
}

func TestGetAllForUser(t *testing.T) {
	repoHomestayMock := new(mocks.HomestayData)
	expectedHomestays := []homestay.Core{
		{
			ID:           1,
			HomestayName: "Lotus",
			Description:  "1 queen size bed",
			Address:      "Bandung",
			Images1:      "http://google.com/images1",
			Images2:      "http://google.com/images2",
			Images3:      "http://google.com/images3",
			CostPerNight: 150000,
		},
		{
			ID:           2,
			HomestayName: "Rose",
			Description:  "2 queen size beds",
			Address:      "Jakarta",
			Images1:      "http://google.com/images4",
			Images2:      "http://google.com/images5",
			Images3:      "http://google.com/images6",
			CostPerNight: 200000,
		},
	}

	repoHomestayMock.On("SelectAllForUser").Return(expectedHomestays, nil)

	srv := New(repoHomestayMock, nil)
	homestays, err := srv.GetAllForUser()

	assert.NoError(t, err)
	assert.Equal(t, expectedHomestays, homestays)
	repoHomestayMock.AssertExpectations(t)

	t.Run("error when getting all homestays for user", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)

		repoHomestayMock.On("SelectAllForUser").Return(nil, errors.New("some error"))

		srv := New(repoHomestayMock, nil)
		homestays, err := srv.GetAllForUser()

		assert.Error(t, err)
		assert.Nil(t, homestays)
		assert.EqualError(t, err, "some error")
		repoHomestayMock.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success get all homestays for host", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)
		repoUserMock := new(mocks.UserData)

		expectedHomestays := []homestay.Core{
			{
				ID:           1,
				HomestayName: "Lotus",
				Description:  "1 queen size bed",
				Address:      "Bandung",
				Images1:      "http://google.com/images1",
				Images2:      "http://google.com/images2",
				Images3:      "http://google.com/images3",
				CostPerNight: 150000,
			},
			{
				ID:           2,
				HomestayName: "Rose",
				Description:  "2 queen size beds",
				Address:      "Jakarta",
				Images1:      "http://google.com/images4",
				Images2:      "http://google.com/images5",
				Images3:      "http://google.com/images6",
				CostPerNight: 200000,
			},
		}

		repoUserMock.On("SelectProfileById", uint(1)).Return(&user.Core{Role: "host"}, nil)
		repoHomestayMock.On("SelectAll", uint(1)).Return(expectedHomestays, nil)

		srv := New(repoHomestayMock, repoUserMock)
		homestays, err := srv.GetAll(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedHomestays, homestays)
		repoHomestayMock.AssertExpectations(t)
		repoUserMock.AssertExpectations(t)
	})

	t.Run("error when getting user profile", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)
		repoUserMock := new(mocks.UserData)

		repoUserMock.On("SelectProfileById", uint(1)).Return(nil, errors.New("user not found"))

		srv := New(repoHomestayMock, repoUserMock)
		homestays, err := srv.GetAll(1)

		assert.Error(t, err)
		assert.Nil(t, homestays)
		assert.EqualError(t, err, "user not found")
		repoHomestayMock.AssertExpectations(t)
		repoUserMock.AssertExpectations(t)
	})

	t.Run("user is not a host", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)
		repoUserMock := new(mocks.UserData)

		repoUserMock.On("SelectProfileById", uint(1)).Return(&user.Core{Role: "user"}, nil)

		srv := New(repoHomestayMock, repoUserMock)
		homestays, err := srv.GetAll(1)

		assert.Error(t, err)
		assert.Nil(t, homestays)
		assert.EqualError(t, err, "you're not host. make a host first")
		repoHomestayMock.AssertExpectations(t)
		repoUserMock.AssertExpectations(t)
	})
}

func TestGetHomestayById(t *testing.T) {
	t.Run("success get homestay by id", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)

		expectedHomestay := homestay.Core{
			ID:           1,
			HomestayName: "Lotus",
			Description:  "1 queen size bed",
			Address:      "Bandung",
			Images1:      "http://google.com/images1",
			Images2:      "http://google.com/images2",
			Images3:      "http://google.com/images3",
			CostPerNight: 150000,
		}

		repoHomestayMock.On("GetHomestayById", uint(1)).Return(expectedHomestay, nil)

		srv := New(repoHomestayMock, nil)
		homestay, err := srv.GetHomestayById(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedHomestay, homestay)
		repoHomestayMock.AssertExpectations(t)
	})

	t.Run("error when getting homestay by id", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)

		repoHomestayMock.On("GetHomestayById", uint(1)).Return(homestay.Core{}, errors.New("homestay not found"))

		srv := New(repoHomestayMock, nil)
		homestay, err := srv.GetHomestayById(1)

		assert.Error(t, err)
		assert.Equal(t, homestay.ID, uint(0))
		assert.EqualError(t, err, "homestay not found")
		repoHomestayMock.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success update homestay", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)
		repoUserMock := new(mocks.UserData)

		id := uint(1)
		idUser := uint(2)
		input := homestay.Core{
			HomestayName: "New name",
			Description:  "Updated description",
			Address:      "Updated address",
			Images1:      "Updated image1",
			Images2:      "Updated image2",
			Images3:      "Updated image3",
			CostPerNight: 200000,
		}

		mockUser := user.Core{Role: "host"}
		repoUserMock.On("SelectProfileById", idUser).Return(&mockUser, nil)

		mockHomestay := homestay.Core{UserID: idUser}
		repoHomestayMock.On("GetUserByHomestayId", id).Return(mockHomestay, nil)
		repoHomestayMock.On("Update", id, input).Return(nil)

		srv := New(repoHomestayMock, repoUserMock)
		err := srv.Update(id, idUser, input)

		assert.NoError(t, err)
		repoHomestayMock.AssertExpectations(t)
		repoUserMock.AssertExpectations(t)
	})

	t.Run("error when user is not a host", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)
		repoUserMock := new(mocks.UserData)

		id := uint(1)
		idUser := uint(2)
		input := homestay.Core{
			HomestayName: "New name",
			Description:  "Updated description",
			Address:      "Updated address",
			Images1:      "Updated image1",
			Images2:      "Updated image2",
			Images3:      "Updated image3",
			CostPerNight: 200000,
		}

		repoUserMock.On("SelectProfileById", idUser).Return(&user.Core{Role: "user"}, nil)

		srv := New(repoHomestayMock, repoUserMock)
		err := srv.Update(id, idUser, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "you're not host. make a host first")
		repoHomestayMock.AssertExpectations(t)
		repoUserMock.AssertExpectations(t)
	})

	t.Run("error when homestay id is not yours", func(t *testing.T) {
		repoHomestayMock := new(mocks.HomestayData)
		repoUserMock := new(mocks.UserData)

		id := uint(1)
		idUser := uint(2)
		input := homestay.Core{
			HomestayName: "New name",
			Description:  "Updated description",
			Address:      "Updated address",
			Images1:      "Updated image1",
			Images2:      "Updated image2",
			Images3:      "Updated image3",
			CostPerNight: 200000,
		}

		mockUser := user.Core{Role: "host"}
		repoUserMock.On("SelectProfileById", idUser).Return(&mockUser, nil)

		mockHomestay := homestay.Core{UserID: idUser}
		repoHomestayMock.On("GetUserByHomestayId", id).Return(mockHomestay, nil)

		repoHomestayMock.On("Update", id, input).Return(errors.New("homestay id is not yours"))

		srv := New(repoHomestayMock, repoUserMock)
		err := srv.Update(id, idUser, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "homestay id is not yours")
		repoHomestayMock.AssertExpectations(t)
		repoUserMock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success delete homestay", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		id := uint(1)
		idUser := uint(2)

		mockUser := user.Core{Role: "host"}
		repoUserMock.On("SelectProfileById", idUser).Return(&mockUser, nil)

		mockHomestay := homestay.Core{UserID: idUser}
		repoHomestayMock.On("GetUserByHomestayId", id).Return(mockHomestay, nil)
		repoHomestayMock.On("Delete", id, idUser).Return(nil)

		srv := New(repoHomestayMock, repoUserMock)
		err := srv.Delete(id, idUser)

		assert.NoError(t, err)
		repoUserMock.AssertExpectations(t)
		repoHomestayMock.AssertExpectations(t)
	})
}

func TestGetMyHomestay(t *testing.T) {
	t.Run("success get my homestay", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		id := uint(1)

		mockUser := &user.Core{Role: "host"}
		repoUserMock.On("SelectProfileById", id).Return(mockUser, nil)

		mockHomestays := []homestay.Core{
			{ID: 1, HomestayName: "Homestay 1"},
			{ID: 2, HomestayName: "Homestay 2"},
		}
		repoHomestayMock.On("GetMyHomestay", id).Return(mockHomestays, nil)

		srv := New(repoHomestayMock, repoUserMock)
		homestays, err := srv.GetMyHomestay(id)

		assert.NoError(t, err)
		assert.NotNil(t, homestays)
		assert.Len(t, homestays, len(mockHomestays))
		for i, h := range homestays {
			assert.Equal(t, mockHomestays[i].ID, h.ID)
			assert.Equal(t, mockHomestays[i].HomestayName, h.HomestayName)
		}
		repoUserMock.AssertExpectations(t)
		repoHomestayMock.AssertExpectations(t)
	})
}

func TestMakeHost(t *testing.T) {
	t.Run("success make host", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		repoHomestayMock := new(mocks.HomestayData)

		id := uint(1)
		input := homestay.Core{
			HomestayName: "Homestay Name",
			Description:  "Description",
			Address:      "Address",
			Images1:      "Image1",
			CostPerNight: 100,
		}

		mockUser := &user.Core{Role: "user"}
		repoUserMock.On("SelectProfileById", id).Return(mockUser, nil)

		repoHomestayMock.On("MakeHost", id, input).Return(nil)

		srv := New(repoHomestayMock, repoUserMock)
		err := srv.MakeHost(id, input)

		assert.NoError(t, err)
		repoUserMock.AssertExpectations(t)
		repoHomestayMock.AssertExpectations(t)
	})

}
