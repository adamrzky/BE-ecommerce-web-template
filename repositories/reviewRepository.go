package repositories

import (
	"BE-ecommerce-web-template/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetMyReview(UserID int) ([]models.Review, error)
	GetReviewByProductID(ProductID int) ([]models.Review, error)
	GetReviewByID(ID int) (models.Review, error)
	Create(review models.Review) (models.Review ,error)
	Update(review models.Review) (models.Review ,error)
	Delete(ID int) error 
}

type reviewRepository struct {
	DB *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{DB: db}
}

func (r *reviewRepository) GetMyReview(UserID int) ([]models.Review, error) {
	var reviews []models.Review
	err := r.DB.Where("user_id = ? ", UserID).Preload("User").Preload("Product").Preload("Transaction").Find(&reviews).Error

	if err != nil {
		return reviews, err
	}

	return reviews, nil
}

func (r *reviewRepository) GetReviewByProductID(ProductID int) ([]models.Review, error) {
	var reviews []models.Review
	err := r.DB.Where("product_id = ? ", ProductID).Preload("User").Preload("Product").Preload("Transaction").Find(&reviews).Error

	if err != nil {
		return reviews, err
	}

	return reviews, nil
}

func (r *reviewRepository) GetReviewByID(ID int) (models.Review, error) {
	var review models.Review
	err := r.DB.Where("id = ? ", ID).Preload("User").Preload("Product").Preload("Transaction").First(&review).Error

	if err != nil {
		return review, err
	}

	return review, nil
}

func (repo *reviewRepository) Create(review models.Review) (models.Review ,error) {
	err := repo.DB.Create(&review).Error

	if err != nil {
		return review, err
	}

	return review, nil
}

func (r *reviewRepository) Update(review models.Review) (models.Review ,error) {
	err :=  r.DB.Model(&models.Review{}).Where("id = ?", review.ID).Updates(review).Error

	if err != nil {
		return review, err
	}

	return review, nil
}

func (r *reviewRepository) Delete(ID int) error {
	return r.DB.Where("id = ?", ID).Delete(&models.Review{}).Error
}