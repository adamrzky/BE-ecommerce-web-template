package repositories

import (
	"BE-ecommerce-web-template/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfileByID(id int) (models.Profile, error)
	GetMyProfile(userID int) (models.Profile, error)
	CreateProfile(profile models.Profile) error
	UpdateProfile(profile models.Profile) error
	DeleteProfile(id int) error
}

type profileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{DB: db}
}

func (r *profileRepository) GetProfileByID(id int) (models.Profile, error) {
	var profile models.Profile
	err := r.DB.Where("id = ?", id).Preload("User").First(&profile).Error
	if err != nil {
		return profile, err
	}
	return profile, err
}

func (r *profileRepository) GetMyProfile(userID int) (models.Profile, error) {
	var profile models.Profile
	err := r.DB.Where("user_id = ?", userID).Preload("User").Find(&profile).Error
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (r *profileRepository) CreateProfile(profile models.Profile) error {
	return r.DB.Create(&profile).Error
}

func (r *profileRepository) UpdateProfile(profile models.Profile) error {
	return r.DB.Save(&profile).Error
}

func (r *profileRepository) DeleteProfile(id int) error {
	return r.DB.Where("id = ?", id).Delete(&models.Profile{}).Error
}
