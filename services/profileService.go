package services

import (
	"BE-ecommerce-web-template/models"
	repository "BE-ecommerce-web-template/repositories"
	"time"
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
}

func (s *ProfileService) Create(userID uint, input ProfileInput) (models.Profile, error) {

	profile := models.Profile{
		Name:    input.Name,
		Gender:  input.Gender,
		Date:    input.Date,
		City:    input.City,
		Address: input.Address,
		Phone:   input.Phone,
		UserID:  userID,
	}
	err := s.ProfileRepo.CreateProfile(profile)
	if err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}

func (s *ProfileService) Update(profileID uint, userID uint, input ProfileInput) (models.Profile, error) {
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
	profile.UserID = userID
	profile.UpdatedAt = time.Now()

	err = s.ProfileRepo.UpdateProfile(profile)
	if err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}

func (s *ProfileService) GetProfileByID(profileID uint) (models.ProfileResponse, error) {
	profile, err := s.ProfileRepo.GetProfileByID(profileID)
	if err != nil {
		return models.ProfileResponse{}, err
	}

	profileResponse := models.ProfileResponse{
		ID:        profile.ID,
		UserID:    profile.UserID,
		Name:      profile.Name,
		Gender:    profile.Gender,
		City:      profile.City,
		Date:      profile.Date,
		Address:   profile.Address,
		Phone:     profile.Phone,
		CreatedAt: profile.CreatedAt,
		UpdatedAt: profile.UpdatedAt,
		User: models.SimpleUserResponse{
			ID:       int(profile.User.ID),
			Email:    profile.User.Email,
			Username: profile.User.Username,
		},
	}

	return profileResponse, nil
}

func (s *ProfileService) GetMyProfile(userID uint) (models.ProfileResponse, error) {
	profile, err := s.ProfileRepo.GetProfileByUserID(userID)
	if err != nil {
		return models.ProfileResponse{}, err
	}

	profileResponse := models.ProfileResponse{
		ID:        profile.ID,
		UserID:    profile.UserID,
		Name:      profile.Name,
		Gender:    profile.Gender,
		City:      profile.City,
		Date:      profile.Date,
		Address:   profile.Address,
		Phone:     profile.Phone,
		CreatedAt: profile.CreatedAt,
		UpdatedAt: profile.UpdatedAt,
		User: models.SimpleUserResponse{
			ID:       int(profile.User.ID),
			Email:    profile.User.Email,
			Username: profile.User.Username,
		},
	}

	return profileResponse, nil
}

func (s *ProfileService) DeleteProfile(id uint) (models.Profile, error) {
	return s.ProfileRepo.DeleteProfile(id)
}
