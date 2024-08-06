package services

import (
	"BE-ecommerce-web-template/models"
	repository "BE-ecommerce-web-template/repositories"
	"BE-ecommerce-web-template/utils/token"
	"time"

	"github.com/gin-gonic/gin"
)

type ProfileService struct {
	ProfileRepo repository.ProfileRepository
}

type ProfileInput struct {
	Name    string    `json:"name"`
	Gender  string    `json:"gender"`
	City    string    `json:"city"`
	Date    time.Time `json:"date"`
	Address string    `json:"address"`
	Phone   string    `json:"phone"`
	UserID  uint      `json:"user_id"`
}

func (s *ProfileService) Create(input ProfileInput) (models.Profile, error) {

	profile := models.Profile{
		Name:    input.Name,
		Gender:  input.Gender,
		Date:    input.Date,
		City:    input.City,
		Address: input.Address,
		Phone:   input.Phone,
		UserID:  input.UserID,
	}
	err := s.ProfileRepo.CreateProfile(profile)
	if err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}

func (s *ProfileService) Update(profileID uint, input ProfileInput) (models.Profile, error) {
	profile, err := s.ProfileRepo.GetProfileByID(profileID)
	if err != nil {
		return models.Profile{}, err
	}

	profile.Name = input.Name
	profile.Gender = input.Gender
	profile.Date = input.Date
	profile.City = input.City
	profile.Address = input.Address
	profile.Phone = input.Phone
	profile.UserID = input.UserID

	err = s.ProfileRepo.UpdateProfile(profile)
	if err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}

func (s *ProfileService) ExtractTokenID(c *gin.Context) (uint, error) {
	return token.ExtractTokenID(c)
}

func (s *ProfileService) GetProfileByID(profileID uint) (models.Profile, error) {
	return s.ProfileRepo.GetProfileByID(profileID)
}
