package models

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ProductsResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Counts  int64       `json:"counts"`
	Data    interface{} `json:"data"`
}
