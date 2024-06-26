package service

import (
	"errors"

	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	"github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user"
)

type homestayService struct {
	homestayData homestay.DataInterface
	userData     user.DataInterface
}

func New(hd homestay.DataInterface, ud user.DataInterface) homestay.ServiceInterface {
	return &homestayService{
		homestayData: hd,
		userData:     ud,
	}
}

// Create implements homestay.ServiceInterface.
func (p *homestayService) Create(input homestay.Core) error {
	if input.HomestayName == "" || input.Address == "" || input.CostPerNight == 0 || input.Description == "" || input.Images1 == "" {
		return errors.New("all list must be filled")
	}
	result, err := p.userData.SelectProfileById(input.UserID)
	if err != nil {
		return err
	}
	if result.Role != "host" {
		return errors.New("you're not host. make a host first")
	}

	err2 := p.homestayData.Insert(input)
	if err2 != nil {
		return err2
	}
	return nil
}

// GetAllForUser implements homestay.ServiceInterface.
func (p *homestayService) GetAllForUser() ([]homestay.Core, error) {
	return p.homestayData.SelectAllForUser()
}

// GetAll implements homestay.ServiceInterface.
func (p *homestayService) GetAll(id uint) ([]homestay.Core, error) {
	result, err := p.userData.SelectProfileById(id)
	if err != nil {
		return nil, err
	}
	if result.Role != "host" {
		return nil, errors.New("you're not host. make a host first")
	}
	return p.homestayData.SelectAll(id)
}

// GetProjectById implements homestay.ServiceInterface.
func (p *homestayService) GetHomestayById(id uint) (input homestay.Core, err error) {
	return p.homestayData.GetHomestayById(id)
}

// Update implements homestay.ServiceInterface.
func (p *homestayService) Update(id uint, idUser uint, input homestay.Core) error {
	result, err := p.userData.SelectProfileById(idUser)
	if err != nil {
		return err
	}
	if result.Role != "host" {
		return errors.New("you're not host. make a host first")
	}

	result2, err2 := p.homestayData.GetUserByHomestayId(id)
	if err2 != nil {
		return err2
	}
	if (result2 == homestay.Core{}) {
		return errors.New("homestay not found")
	}
	if result2.UserID != idUser {
		return errors.New("homestay id is not yours")
	}
	return p.homestayData.Update(id, input)
}

// Delete implements homestay.ServiceInterface.
func (p *homestayService) Delete(id uint, idUser uint) error {
	result, err := p.userData.SelectProfileById(idUser)
	if err != nil {
		return err
	}
	if result.Role != "host" {
		return errors.New("you're not host. make a host first")
	}
	result2, err2 := p.homestayData.GetUserByHomestayId(id)
	if err2 != nil {
		return err2
	}
	if result2.UserID != idUser {
		return errors.New("homestay id is not yours")
	}
	return p.homestayData.Delete(id, idUser)
}

// GetMyHomestay implements homestay.ServiceInterface.
func (p *homestayService) GetMyHomestay(id uint) ([]homestay.Core, error) {
	result, err := p.userData.SelectProfileById(id)
	if err != nil {
		return nil, err
	}
	if result.Role != "host" {
		return nil, errors.New("you're not host. make a host first")
	}
	return p.homestayData.GetMyHomestay(id)
}

// MakeHost implements homestay.ServiceInterface.
func (p *homestayService) MakeHost(id uint, input homestay.Core) error {
	result, err := p.userData.SelectProfileById(id)
	if err != nil {
		return err
	}
	if result.Role == "host" {
		return errors.New("you're already host")
	}

	if input.HomestayName == "" || input.Address == "" || input.CostPerNight == 0 || input.Description == "" || input.Images1 == "" {
		return errors.New("all list must be filled")
	}

	return p.homestayData.MakeHost(id, input)
}
