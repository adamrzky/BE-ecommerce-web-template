package services

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/repositories"
	"time"
)

type ProfileService interface {
	GetProfileByID(id int) (models.ProfileResponse, error)
	Update(profileID int, userID int, input ProfileInput) (models.Profile, error)
	Create(userID int, input ProfileInput) (models.Profile, error)
	DeleteProfile(id int) error
	GetMyProfile(userID int) (models.ProfileResponse, error)
}

type profileService struct {
	ProfileRepo repositories.ProfileRepository
}

func NewProfileService(repo repositories.ProfileRepository) ProfileService {
	return &profileService{ProfileRepo: repo}
}

type ProfileInput struct {
	Name    string    `json:"name"`
	Gender  string    `json:"gender"`
	City    string    `json:"city"`
	Date    time.Time `json:"date"`
	Address string    `json:"address"`
	Phone   string    `json:"phone"`
}

func (s *profileService) Create(userID int, input ProfileInput) (models.Profile, error) {
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
	return profile, err
}

func (s *profileService) Update(profileID int, userID int, input ProfileInput) (models.Profile, error) {
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
	return profile, err
}

func (s *profileService) GetProfileByID(profileID int) (models.ProfileResponse, error) {
	profile, err := s.ProfileRepo.GetProfileByID(profileID)
	if err != nil {
		return models.ProfileResponse{}, err
	}

	profilesResponse := models.ProfileResponse{
		ID:        int(profile.ID),
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

	return profilesResponse, nil
}

func (s *profileService) GetMyProfile(userID int) (models.ProfileResponse, error) {
	profile, err := s.ProfileRepo.GetMyProfile(userID)
	if err != nil {
		return models.ProfileResponse{}, err
	}

	profilesResponse := models.ProfileResponse{
		ID:        int(profile.ID),
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

	return profilesResponse, nil
}

func (s *profileService) DeleteProfile(id int) error {
	return s.ProfileRepo.DeleteProfile(id)
}
