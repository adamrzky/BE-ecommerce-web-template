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
	GetLikesByUserID(userID uint) ([]models.UserProductLikes, error)
	CompositeLikeExist(userID, prpoductID uint) (bool, error)
	CountProducts(params models.ProductQueryParam) (int64, error)
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

func (repo *productRepository) CountProducts(params models.ProductQueryParam) (int64, error) {
	var count int64
	query := repo.db.Model(&models.Product{}).Preload("Category").Preload("ProductProps")

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

	err := query.Count(&count).Error
	return count, err
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
	// return repo.db.Create(&product).Error
	return repo.db.Transaction(func(tx *gorm.DB) error {
		// makes product creation below to return id so i can use for assign to product props
		if err := tx.Create(&product).Error; err != nil {
			return err
		}
		product.ProductProps.ProductID = product.ID
		return tx.Create(&product.ProductProps).Error
	})
}

func (repo *productRepository) Update(product *models.Product, id uint) error {
	return repo.db.Model(&models.Product{}).Where("ID = ?", id).Updates(product).Error
}

func (repo *productRepository) Delete(id uint) error {
	return repo.db.Exec("CALL DeleteProduct(?)", id).Error
}

func (repo *productRepository) PostLike(userID, productID uint) error {
	return repo.db.Exec("CALL LikeProduct(?, ?)", userID, productID).Error
}

func (repo *productRepository) DeleteLike(userID, productID uint) error {
	return repo.db.Exec("CALL DislikeProduct(?, ?)", userID, productID).Error
}

func (repo *productRepository) GetLikesByUserID(userID uint) ([]models.UserProductLikes, error) {
	var likes []models.UserProductLikes
	err := repo.db.Preload("Product").Where("user_id = ?", userID).Find(&likes).Error
	return likes, err
}

func (repo *productRepository) CompositeLikeExist(userID, productID uint) (bool, error) {
	var count int64
	err := repo.db.Model(&models.UserProductLikes{}).Where("user_id = ? AND product_id = ?", userID, productID).Count(&count).Error
	return count > 0, err
}
