package services

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/repositories"
)

type ReviewService interface {
	GetMyReview(UserID int) ([]models.ReviewResponse, error)
	GetReviewByProductID(ProductID int) ([]models.ReviewResponse, error)
	GetReviewByID(ID int) (models.ReviewResponse, error)
	CreateReview(input models.ReviewInput, UserID int) (models.Review,error)
	UpdateReview(input models.ReviewInput,  UserID int, reviewId int) (models.Review,error)
	DeleteReview(reviewID int) error
}

type reviewService struct {
	repo repositories.ReviewRepository
}

func NewReviewService(repo repositories.ReviewRepository) ReviewService {
	return &reviewService{repo}
}

func(s *reviewService) GetMyReview(UserID int) ([]models.ReviewResponse, error) {
	reviews, err := s.repo.GetMyReview(UserID)
	if err != nil {
		return nil, err
	}

	var reviewsResponses []models.ReviewResponse
	for _, review := range reviews {
		reviewData := models.ReviewResponse{
			ID: int(review.ID),
			UserID: review.UserID,
			ProductID: review.ProductID,
			TransactionID: review.TransactionID,
			Content: review.Content,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
			User: models.SimpleUserResponse{
				ID: int(review.User.ID),
				Email: review.User.Email,
				Username: review.User.Username,
			},
			Product: models.SimpleProductResponse{
				ID: int(review.Product.ID),
				Name: review.Product.Name,
			},
			Transaction: models.SimpleTrxResponse{
				ID: int(review.Transaction.ID),
				TransactionID: review.Transaction.TRX_ID,
				Status: review.Transaction.STATUS,
			},
		}
		reviewsResponses = append(reviewsResponses, reviewData)
	}

	return reviewsResponses, nil
}

func(s *reviewService) GetReviewByProductID(ProductID int) ([]models.ReviewResponse, error) {
	reviews, err := s.repo.GetReviewByProductID(ProductID)
	if err != nil {
		return nil, err
	}

	var reviewsResponses []models.ReviewResponse
	for _, review := range reviews {
		reviewsResponse := models.ReviewResponse{
			ID: int(review.ID),
			UserID: review.UserID,
			ProductID: review.ProductID,
			TransactionID: review.TransactionID,
			Content: review.Content,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
			User: models.SimpleUserResponse{
				ID: int(review.User.ID),
				Email: review.User.Email,
				Username: review.User.Username,
			},
			Product: models.SimpleProductResponse{
				ID: int(review.Product.ID),
				Name: review.Product.Name,
			},
			Transaction: models.SimpleTrxResponse{
				ID: int(review.Transaction.ID),
				TransactionID: review.Transaction.TRX_ID,
				Status: review.Transaction.STATUS,
			},
		}
		reviewsResponses = append(reviewsResponses, reviewsResponse)
	}

	return reviewsResponses, nil
}

func(s *reviewService) GetReviewByID(ID int) (models.ReviewResponse, error) {
	review, err := s.repo.GetReviewByID(ID)
	if err != nil {
		return models.ReviewResponse{}, err
	}

	reviewsResponse := models.ReviewResponse{
		ID: int(review.ID),
		UserID: review.UserID,
		ProductID: review.ProductID,
		TransactionID: review.TransactionID,
		Content: review.Content,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
		User: models.SimpleUserResponse{
			ID: int(review.User.ID),
			Email: review.User.Email,
			Username: review.User.Username,
		},
		Product: models.SimpleProductResponse{
			ID: int(review.Product.ID),
			Name: review.Product.Name,
		},
		Transaction: models.SimpleTrxResponse{
			ID: int(review.Transaction.ID),
			TransactionID: review.Transaction.TRX_ID,
			Status: review.Transaction.STATUS,
		},
	}

	return reviewsResponse, nil
}

func (s *reviewService) CreateReview(input models.ReviewInput, UserID int) (models.Review,error) {
	newReview := models.Review{
		UserID: UserID,
		ProductID: input.ProductID,
		TransactionID: input.TransactionID,
		Content: input.Content,
	}

	review, err := s.repo.Create(newReview)
	if err != nil {
		return review, err
	}

	return review, nil
}

func (s *reviewService) UpdateReview(input models.ReviewInput, UserID int, reviewID int) (models.Review,error) {
	updatedReview := models.Review{
		ID: uint(reviewID),
		UserID: UserID,
		ProductID: input.ProductID,
		TransactionID: input.TransactionID,
		Content: input.Content,
	}

	review, err := s.repo.Update(updatedReview)
	if err != nil {
		return review, err
	}

	return review, nil
}

func (s *reviewService) DeleteReview(reviewID int) error {
	return s.repo.Delete(reviewID)
}