package service

import (
	"errors"
	homestay "myTaskApp/features/homestays"
	"myTaskApp/features/user"
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
	result, err := p.userData.SelectProfileById(input.UserID)
	if err != nil {
		return err
	}
	if result.Role != "hoster" {
		return errors.New("you're not hoster. switch your role first")
	}
	if result.ID != input.UserID {
		return errors.New("user not found")
	}
	if input.HomestayName == "" {
		return errors.New("homestay name must be filled")
	}

	err2 := p.homestayData.Insert(input)
	if err2 != nil {
		return err2
	}
	return nil
}

// GetAllForUser implements homestay.ServiceInterface.
func (p *homestayService) GetAllForUser(id uint) ([]homestay.Core, error) {
	return p.homestayData.SelectAllForUser(id)
}

// GetAll implements homestay.ServiceInterface.
func (p *homestayService) GetAll(id uint) ([]homestay.Core, error) {
	result, err := p.userData.SelectProfileById(id)
	if err != nil {
		return nil, err
	}
	if result.Role != "hoster" {
		return nil, errors.New("you're not hoster. switch your role first")
	}
	return p.homestayData.SelectAll(id)
}

// GetProjectById implements homestay.ServiceInterface.
func (p *homestayService) GetHomestayById(id uint, idUser uint) (input homestay.Core, err error) {
	result, err := p.userData.SelectProfileById(input.UserID)
	if err != nil {
		return homestay.Core{}, err
	}
	if result.Role != "hoster" {
		return homestay.Core{}, errors.New("you're not hoster. switch your role first")
	}
	result2, err2 := p.homestayData.GetUserByHomestayId(id)
	if err2 != nil {
		return homestay.Core{}, err2
	}
	if result2.UserID != idUser {
		return homestay.Core{}, errors.New("homestay id is not yours")
	}
	return p.homestayData.GetHomestayById(id)
}

// Update implements homestay.ServiceInterface.
func (p *homestayService) Update(id uint, idUser uint, input homestay.Core) error {
	result, err := p.userData.SelectProfileById(input.UserID)
	if err != nil {
		return err
	}
	if result.Role != "hoster" {
		return errors.New("you're not hoster. switch your role first")
	}
	result2, err2 := p.homestayData.GetUserByHomestayId(id)
	if err2 != nil {
		return err2
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
	if result.Role != "hoster" {
		return errors.New("you're not hoster. switch your role first")
	}
	result2, err2 := p.homestayData.GetUserByHomestayId(id)
	if err2 != nil {
		return err2
	}
	if result2.UserID != idUser {
		return errors.New("homestay id is not yours")
	}
	return p.homestayData.Delete(id)
}
