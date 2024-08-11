package models

type (
	ProductResponse struct {
		ID           uint                 `json:"id"`
		Name         string               `json:"name"`
		Description  string               `json:"description"`
		Price        float32              `json:"price"`
		Slug         string               `json:"slug"`
		ImageURL     string               `json:"image_url,omitempty"`
		Status       int                  `json:"status,omitempty"`
		CategoryID   uint                 `json:"category_id,omitempty"`
		Category     CategoryResponse     `json:"category,omitempty"`
		ProductProps ProductPropsResponse `json:"product_props,omitempty"`
	}

	ProductRequest struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
		ImageURL    string  `json:"image_url"`
		Status      int     `json:"status"`
		CategoryID  uint    `json:"category_id"`
	}

	ProductQueryParam struct {
		ProductName string  `form:"productName"`
		Category    uint    `form:"category"`
		MinPrice    float32 `form:"minPrice"`
		MaxPrice    float32 `form:"maxPrice"`
		Limit       int     `form:"limit" validate:"required,min=1"`
		Offset      int     `form:"offset"`
	}
)
