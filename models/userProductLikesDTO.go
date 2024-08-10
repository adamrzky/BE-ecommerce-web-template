package models

type (
	UserProductLikesResponse struct {
		ID        uint            `json:"id"`
		UserID    uint            `json:"user_id"`
		ProductID uint            `json:"product_id"`
		Product   ProductResponse `json:"product"`
	}

	UserProductLikesRequest struct {
		ProductID uint `json:"product_id"`
	}
)
