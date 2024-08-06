package models

type (
	ProductResponse struct {
		ID         uint             `json:"ID"`
		Name       string           `json:"name"`
		Price      float32          `json:"price"`
		Slug       string           `json:"slug"`
		ImageURL   string           `json:"image_url,omitempty"`
		Status     int              `json:"status,omitempty"`
		CategoryID uint             `json:"category_id,omitempty"`
		Category   CategoryResponse `json:"category,omitempty"`
	}

	ProductRequest struct {
		Name       string  `json:"name" validate:"required"`
		Price      float32 `json:"price" validate:"required"`
		Slug       string  `json:"slug" validate:"required"`
		ImageURL   string  `json:"image_url"`
		Status     int     `json:"status"`
		CategoryID uint    `json:"category_id"`
	}

	ProductQueryParam struct {
		MinPrice float32 `form:"min_price"`
		MaxPrice float32 `form:"max_price"`
		Limit    int     `form:"limit" validate:"required,min=1"`
		Offset   int     `form:"offset"`
	}
)
