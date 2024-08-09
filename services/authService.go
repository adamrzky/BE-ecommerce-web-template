package services

import (
	"BE-ecommerce-web-template/models"
	repository "BE-ecommerce-web-template/repositories"
	"BE-ecommerce-web-template/utils/token"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo repository.UserRepository
}

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *AuthService) Register(input RegisterInput) (models.User, error) {
	// Validasi email
	if !isValidEmail(input.Email) {
		return models.User{}, errors.New("invalid email address")
	}

	// Validasi password
	if err := validatePassword(input.Password); err != nil {
		return models.User{}, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Email:    input.Email,
		RoleID:   1,
	}

	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}

func (s *AuthService) Login(input LoginInput) (string, models.User, error) {
	user, err := s.UserRepo.GetUserByUsername(input.Username)
	if err != nil {
		return "", models.User{}, errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", models.User{}, errors.New("invalid username or password")
	}

	// Generate token with id, name and role to extract with
	tokenStr, err := token.GenerateToken(user.ID, user.Username, user.Role.Name)
	if err != nil {
		return "", models.User{}, err
	}

	return tokenStr, user, nil
}

func isValidEmail(email string) bool {
	return email != ""
}

func (s *AuthService) ExtractTokenID(c *gin.Context) (uint, error) {
	return token.ExtractTokenID(c)
}

func (s *AuthService) GetUserByID(userID uint) (models.User, error) {
	return s.UserRepo.GetUserByID(userID)
}

type ChangePasswordInput struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

func (s *AuthService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("old password is incorrect")
	}

	if err := validatePassword(newPassword); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	err = s.UserRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
