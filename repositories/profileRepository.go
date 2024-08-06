package repositories

import (
	"BE-ecommerce-web-template/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfileByID(id uint) (models.Profile, error)
	UpdateProfile(profile models.Profile) error
	CreateProfile(profile models.Profile) error
}

type profileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{DB: db}
}

func (r *profileRepository) GetProfileByID(id uint) (models.Profile, error) {
	var profile models.Profile
	if err := r.DB.Preload("User").First(&profile, id).Error; err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}

func (r *profileRepository) CreateProfile(profile models.Profile) error {
	return r.DB.Create(&profile).Error
}

func (r *profileRepository) UpdateProfile(profile models.Profile) error {
	return r.DB.Save(&profile).Error
}
