package repositories

import (
	"BE-ecommerce-web-template/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfileByID(id uint) (models.Profile, error)
	UpdateProfile(profile models.Profile) error
	CreateProfile(profile models.Profile) error
	DeleteProfile(id uint) (models.Profile, error)
	GetProfileByUserID(userID uint) (models.Profile, error)
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

func (r *profileRepository) GetProfileByUserID(userID uint) (models.Profile, error) {
	var profile models.Profile
	if err := r.DB.Preload("User").Where("user_id = ?", userID).First(&profile).Error; err != nil {
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

func (r *profileRepository) DeleteProfile(id uint) (models.Profile, error) {
	var profile models.Profile
	if err := r.DB.First(&profile, id).Error; err != nil {
		return models.Profile{}, err
	}

	if err := r.DB.Delete(&profile).Error; err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}
