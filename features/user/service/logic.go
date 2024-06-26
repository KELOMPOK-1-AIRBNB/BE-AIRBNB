package service

import (
	"errors"

	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/app/middlewares"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user"
	encrypts "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/utils"
)

type userService struct {
	userData    user.DataInterface
	hashService encrypts.HashInterface
}

func New(ud user.DataInterface, hash encrypts.HashInterface) user.ServiceInterface {
	return &userService{
		userData:    ud,
		hashService: hash,
	}
}

func (u *userService) UpdateProfilePicture(id uint, input user.Core) error {
	result, err := u.userData.SelectProfileById(id)
	if err != nil {
		return errors.New("user not found. you must login first")
	}

	if result.DeleteAt.IsZero() {
		return u.userData.UpdateProfilePicture(id, input)
	} else {
		return errors.New("user not found. you must login first")
	}
}

// Create implements user.ServiceInterface.
func (u *userService) Create(input user.Core) error {
	if input.Name == "" || input.Email == "" || input.Password == "" || input.Phone == "" {
		return errors.New("[validation] nama/email/password/phone tidak boleh kosong")
	}
	result, errHash := u.hashService.HashPassword(input.Password)
	if errHash != nil {
		return errHash
	}
	input.Password = result

	err := u.userData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements user.ServiceInterface.
func (u *userService) GetProfileUser(id uint) (*user.Core, error) {
	if id <= 0 {
		return nil, errors.New("id not valid")
	}
	return u.userData.SelectProfileById(id)
}

// Delete implements user.ServiceInterface.
func (u *userService) Delete(id uint) error {
	result, err := u.userData.SelectProfileById(id)
	if err != nil {
		return errors.New("user not found. you must login first")
	}
	if result.DeleteAt.IsZero() {
		return u.userData.Delete(id)
	} else {
		return errors.New("user not found. you must login first")
	}
}

// Update implements user.ServiceInterface.
func (u *userService) Update(id uint, input user.Core) error {
	result, err := u.userData.SelectProfileById(id)
	if err != nil {
		return errors.New("user not found. you must login first")
	}

	result2, errHash := u.hashService.HashPassword(input.Password)
	if errHash != nil {
		return errHash
	}
	if input.Password != "" {
		input.Password = result2
	}

	if result.DeleteAt.IsZero() {
		return u.userData.Update(id, input)
	} else {
		return errors.New("user not found. you must login first")
	}
}

// Login implements user.ServiceInterface.
func (u *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	data, err = u.userData.Login(email)
	if err != nil {
		return nil, "", err
	}

	isLoginValid := u.hashService.CheckPasswordHash(data.Password, password)
	if !isLoginValid {
		return nil, "", errors.New("[validation] password tidak sesuai")
	}
	token, errJWT := middlewares.CreateToken(int(data.ID))
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil
}
