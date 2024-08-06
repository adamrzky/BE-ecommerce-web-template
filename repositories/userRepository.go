package repositories

import (
	"BE-ecommerce-web-template/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id uint) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	UpdateUser(user models.User) error
	CreateUser(user models.User) error
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
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := r.DB.Preload("Role").Where("USERNAME = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user models.User) error {
	return r.DB.Create(&user).Error
}

func (r *userRepository) UpdateUser(user models.User) error {
	return r.DB.Save(&user).Error
}
