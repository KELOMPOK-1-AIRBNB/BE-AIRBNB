package data

import (
	homestay "github.com/KELOMPOK-1-AIRBNB/BE-AIRBNB/features/homestays"

	"gorm.io/gorm"
)

type homestayQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) homestay.DataInterface {
	return &homestayQuery{
		db: db,
	}
}

// Insert implements homestay.DataInterface.
func (p *homestayQuery) Insert(input homestay.Core) error {
	var homestayGorm Homestay

	homestayGorm = Homestay{
		Model:        gorm.Model{},
		UserID:       input.UserID,
		HomestayName: input.HomestayName,
		Description:  input.Description,
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
func (p *homestayQuery) Delete(id uint) error {
	tx2 := p.db.Delete(&Homestay{}, id)
	if tx2.Error != nil {
		return tx2.Error
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
		})
	}

	return allHomestayCore, nil
}
