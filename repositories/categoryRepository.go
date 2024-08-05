package repository

import (
	"BE-ecommerce-web-template/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id uint) (models.Category, error)
	Post(category models.Category) error
	Update(category *models.Category, id uint) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (repo *categoryRepository) GetAll() ([]models.Category, error) {
	var category []models.Category
	err := repo.db.Find(&category).Error
	return category, err
}

func (repo *categoryRepository) GetByID(id uint) (models.Category, error) {
	var category models.Category
	err := repo.db.Where("id = ?", id).First(&category).Error
	return category, err
}

func (repo *categoryRepository) Post(category models.Category) error {
	return repo.db.Create(&category).Error
}

func (repo *categoryRepository) Update(category *models.Category, id uint) error {
	return repo.db.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error
}

func (repo *categoryRepository) Delete(id uint) error {
	return repo.db.Where("id = ?", id).Delete(&models.Category{}).Error
}
