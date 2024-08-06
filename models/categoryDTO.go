package models

type (
	CategoryResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	CategoryRequest struct {
		Name string `json:"name" validate:"required"`
	}
)
