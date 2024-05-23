package data

import (
	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"
	userInterface "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/user"

	"gorm.io/gorm"
)

type homestayQuery struct {
	db   *gorm.DB
	user userInterface.DataInterface
}

func New(db *gorm.DB, user userInterface.DataInterface) homestay.DataInterface {
	return &homestayQuery{
		db:   db,
		user: user,
	}
}

// Insert implements homestay.DataInterface.
func (p *homestayQuery) Insert(input homestay.Core) error {
	homestayGorm := Homestay{
		Model:         gorm.Model{},
		UserID:        input.UserID,
		HomestayName:  input.HomestayName,
		Address:       input.Address,
		Images1:       input.Images1,
		Images2:       input.Images2,
		Images3:       input.Images3,
		Description:   input.Description,
		PricePerNight: input.CostPerNight,
	}

	tx := p.db.Create(&homestayGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectAllForUser implements homestay.DataInterface.
func (p *homestayQuery) SelectAllForUser() ([]homestay.Core, error) {
	var allHomestay []Homestay
	tx := p.db.Find(&allHomestay)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var allHomestayCore []homestay.Core
	for _, v := range allHomestay {
		allHomestayCore = append(allHomestayCore, homestay.Core{
			ID:           v.ID,
			UserID:       v.UserID,
			HomestayName: v.HomestayName,
			Address:      v.Address,
			Description:  v.Description,
			CostPerNight: v.PricePerNight,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}

	return allHomestayCore, nil
}

// SelectAll implements homestay.DataInterface.
func (p *homestayQuery) SelectAll(id uint) ([]homestay.Core, error) {
	var allHomestay []Homestay
	tx := p.db.Where("user_id != ?", id).Find(&allHomestay)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var allHomestayCore []homestay.Core
	for _, v := range allHomestay {
		allHomestayCore = append(allHomestayCore, homestay.Core{
			ID:           v.ID,
			UserID:       v.UserID,
			HomestayName: v.HomestayName,
			Address:      v.Address,
			Description:  v.Description,
			CostPerNight: v.PricePerNight,
		})
	}

	return allHomestayCore, nil
}

// GetHomestayById implements homestay.DataInterface.
func (p *homestayQuery) GetHomestayById(id uint) (homestay.Core, error) {
	var homestayId Homestay
	tx := p.db.First(&homestayId, id)
	if tx.Error != nil {
		return homestay.Core{}, tx.Error
	}

	homestayIdCore := homestay.Core{
		ID:           id,
		UserID:       homestayId.UserID,
		HomestayName: homestayId.HomestayName,
		Address:      homestayId.Address,
		Description:  homestayId.Description,
		CostPerNight: homestayId.PricePerNight,
	}

	return homestayIdCore, nil

}

// GetHomestayByUserId implements homestay.DataInterface.
func (p *homestayQuery) GetHomestayByUserId(id uint) ([]homestay.Core, error) {
	var allHomestayUser []Homestay
	tx := p.db.Where("user_id = ?", id).Find(&allHomestayUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var allHomestayUserCore []homestay.Core
	for _, v := range allHomestayUser {
		allHomestayUserCore = append(allHomestayUserCore, homestay.Core{
			ID:           v.ID,
			UserID:       v.UserID,
			HomestayName: v.HomestayName,
			Address:      v.Address,
			Description:  v.Description,
			CostPerNight: v.PricePerNight,
		})
	}

	return allHomestayUserCore, nil
}

// Update implements homestay.DataInterface.
func (p *homestayQuery) Update(id uint, input homestay.Core) error {
	updateList := Homestay{
		HomestayName:  input.HomestayName,
		Address:       input.Address,
		Images1:       input.Images1,
		Images2:       input.Images2,
		Images3:       input.Images3,
		Description:   input.Description,
		PricePerNight: input.CostPerNight,
	}
	tx2 := p.db.Model(&Homestay{}).Where("id = ?", id).Updates(updateList)
	if tx2.Error != nil {
		return tx2.Error
	}
	return nil
}

// Delete implements homestay.DataInterface.
func (p *homestayQuery) Delete(id uint, idUser uint) error {
	tx1 := p.db.Delete(&Homestay{}, id)
	if tx1.Error != nil {
		return tx1.Error
	}

	result, err := p.GetHomestayByUserId(idUser)
	if err != nil {
		return err
	}

	role := userInterface.Core{
		Role: "user",
	}

	if len(result) == 0 {
		if err := p.user.UpdateRole(idUser, role); err != nil {
			return err
		}
	}

	return nil
}

// GetUserByHomestayId implements homestay.DataInterface.
func (p *homestayQuery) GetUserByHomestayId(id uint) (homestay.Core, error) {
	var homestayUserId Homestay
	tx := p.db.First(&homestayUserId, id)
	if tx.Error != nil {
		return homestay.Core{}, tx.Error
	}
	projectIdCore := homestay.Core{
		ID:     id,
		UserID: homestayUserId.UserID,
	}

	return projectIdCore, nil
}

// GetMyHomestay implements homestay.DataInterface.
func (p *homestayQuery) GetMyHomestay(id uint) ([]homestay.Core, error) {
	var allHomestay []Homestay
	tx := p.db.Where("user_id = ?", id).Find(&allHomestay)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var allHomestayCore []homestay.Core
	for _, v := range allHomestay {
		allHomestayCore = append(allHomestayCore, homestay.Core{
			ID:           v.ID,
			UserID:       v.UserID,
			HomestayName: v.HomestayName,
			Address:      v.Address,
			Description:  v.Description,
			CostPerNight: v.PricePerNight,
			Images1:      v.Images1,
			Images2:      v.Images2,
			Images3:      v.Images3,
		})
	}

	return allHomestayCore, nil
}

// MakeHost implements homestay.DataInterface.
func (p *homestayQuery) MakeHost(id uint, input homestay.Core) error {
	role := userInterface.Core{
		Role: "host",
	}
	p.user.UpdateRole(id, role)

	err := p.Insert(input)
	if err != nil {
		return err
	}

	return nil
}
