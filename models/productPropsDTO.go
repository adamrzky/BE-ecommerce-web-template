package models

type (
	ProductPropsResponse struct {
		ID            uint    `json:"id"`
		TotalLikes    uint    `json:"total_likes"`
		TotalReviews  uint    `json:"total_reviews"`
		AverageRating float32 `json:"average_rating"`
		ProductID     uint    `json:"product_id"`
	}

	ProductPropsRequest struct {
		TotalLikes    uint    `json:"total_likes"`
		TotalReviews  uint    `json:"total_reviews"`
		AverageRating float32 `json:"average_rating"`
		ProductID     uint    `json:"product_id"`
	}
)
