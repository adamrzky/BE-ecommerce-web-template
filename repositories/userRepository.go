package repositories

import (
	"BE-ecommerce-web-template/models"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	GetUserByID(id uint) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByUsername(username string) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := r.DB.Preload("Role").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, ErrUserNotFound
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Preload("Role").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := r.DB.Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, ErrUserNotFound
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user models.User) error {
	return r.DB.Create(&user).Error
}

func (r *userRepository) UpdateUser(user models.User) error {
	query := `UPDATE USER SET USERNAME = ?, PASSWORD = ?, EMAIL = ?, ROLE_ID = ? WHERE ID = ?`

	result := r.DB.Exec(query, user.Username, user.Password, user.Email, user.RoleID, user.ID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows affected, possible issue with ID")
	}

	return nil
}

func (r *userRepository) DeleteUser(id uint) error {
	if err := r.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
