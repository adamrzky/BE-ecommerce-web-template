package repositories

import (
	"BE-ecommerce-web-template/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll(params models.ProductQueryParam) ([]models.Product, error)
	GetByID(id uint) (models.Product, error)
	GetBySlug(slug string) (models.Product, error)
	Post(product models.Product) error
	Update(product *models.Product, id uint) error
	Delete(id uint) error
	PostLike(userID, productID uint) error
	DeleteLike(userID, productID uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (repo *productRepository) GetAll(params models.ProductQueryParam) ([]models.Product, error) {
	var products []models.Product
	query := repo.db.Preload("Category").Preload("ProductProps")

	// Name filter
	if params.ProductName != "" {
		query = query.Where("NAME LIKE ?", "%"+params.ProductName+"%")
	}

	// Category filter
	if params.Category != 0 {
		query = query.Where("CATEGORY_ID = ?", params.Category)
	}

	// Price filter
	if params.MinPrice > 0 {
		query = query.Where("PRICE >= ?", params.MinPrice)
	}
	if params.MaxPrice > 0 {
		query = query.Where("PRICE <= ?", params.MaxPrice)
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
	err := repo.db.Preload("Category").Preload("ProductProps").Where("ID = ?", id).First(&product).Error
	return product, err
}

func (repo *productRepository) GetBySlug(slug string) (models.Product, error) {
	var product models.Product
	err := repo.db.Preload("Category").Preload("ProductProps").Where("slug = ?", slug).First(&product).Error
	return product, err
}

func (repo *productRepository) Post(product models.Product) error {
	return repo.db.Create(&product).Error
}

func (repo *productRepository) Update(product *models.Product, id uint) error {
	return repo.db.Model(&models.Product{}).Where("ID = ?", id).Updates(product).Error
}

func (repo *productRepository) Delete(id uint) error {
	return repo.db.Where("ID = ?", id).Delete(&models.Product{}).Error
}

func (repo *productRepository) PostLike(userID, productID uint) error {
	return repo.db.Exec("CALL LikeProduct(?, ?)", userID, productID).Error
}

func (repo *productRepository) DeleteLike(userID, productID uint) error {
	return repo.db.Exec("CALL DislikeProduct(?, ?)", userID, productID).Error
}
