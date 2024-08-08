package services

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserByID(id uint) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	CreateUser(input CreateUserInput) (models.User, error)
	UpdateUser(id uint, input UpdateUserInput) (models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	UserRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{UserRepo: repo}
}

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleID   uint   `json:"role_id"`
}

type UpdateUserInput struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Email    *string `json:"email,omitempty"`
	RoleID   *uint   `json:"role_id,omitempty"`
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *userService) GetUserByUsername(username string) (models.User, error) {
	return s.UserRepo.GetUserByUsername(username)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *userService) CreateUser(input CreateUserInput) (models.User, error) {
	if !isValidEmail(input.Email) {
		return models.User{}, errors.New("invalid email address")
	}

	if err := validatePassword(input.Password); err != nil {
		return models.User{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Email:    input.Email,
		RoleID:   input.RoleID,
	}

	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) UpdateUser(id uint, input UpdateUserInput) (models.User, error) {
	user, err := s.UserRepo.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}

	if input.Username != nil {
		user.Username = *input.Username
	}
	if input.Password != nil {
		if err := validatePassword(*input.Password); err != nil {
			return models.User{}, err
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, err
		}
		user.Password = string(hashedPassword)
	}
	if input.Email != nil {
		if !isValidEmail(*input.Email) {
			return models.User{}, errors.New("invalid email address")
		}
		user.Email = *input.Email
	}
	if input.RoleID != nil {
		user.RoleID = *input.RoleID
	}

	err = s.UserRepo.UpdateUser(user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.UserRepo.DeleteUser(id)
}
