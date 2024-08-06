package repositories

import (
	"BE-ecommerce-web-template/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll(params models.ProductQueryParam) ([]models.Product, error)
	GetByID(id uint) (models.Product, error)
	Post(product models.Product) error
	Update(product *models.Product, id uint) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (repo *productRepository) GetAll(params models.ProductQueryParam) ([]models.Product, error) {
	var products []models.Product
	query := repo.db.Preload("Category")

	// Price filter
	if params.MinPrice > 0 {
		query = query.Where("price >= ?", params.MinPrice)
	}
	if params.MaxPrice > 0 {
		query = query.Where("price <= ?", params.MaxPrice)
	}

	// Pagination
	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}
	query = query.Offset(params.Offset)

	err := query.Find(&products).Error
	return products, err
}

func (repo *productRepository) GetByID(id uint) (models.Product, error) {
	var product models.Product
	err := repo.db.Preload("Category").Where("id = ?", id).First(&product).Error
	return product, err
}

func (repo *productRepository) Post(product models.Product) error {
	return repo.db.Create(&product).Error
}

func (repo *productRepository) Update(product *models.Product, id uint) error {
	return repo.db.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (repo *productRepository) Delete(id uint) error {
	return repo.db.Where("id = ?", id).Delete(&models.Product{}).Error
}
