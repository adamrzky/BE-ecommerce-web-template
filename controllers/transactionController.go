package controllers

import (
	"net/http"
	"strconv"

	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/services"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

// GetTransactionByID godoc
// @Summary Retrieve a transaction
// @Description get transaction by ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /transactions/{id} [get]
func (tc *TransactionController) GetTransactionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID format",
		})
		return
	}
	transaction, err := tc.service.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Transaction retrieved successfully",
		Data:    transaction,
	})
}

// CreateTransaction godoc
// @Summary Create a transaction
// @Description create a new transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param transaction body models.Transaction true "Create Transaction"
// @Success 201 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /transactions [post]
func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid data provided",
		})
		return
	}
	if err := tc.service.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transaction",
		})
		return
	}
	c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "Transaction created successfully",
		Data:    transaction,
	})
}

// UpdateTransaction godoc
// @Summary Update a transaction
// @Description update a transaction by ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Param transaction body models.Transaction true "Update Transaction"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /transactions/{id} [put]
func (tc *TransactionController) UpdateTransaction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID format",
		})
		return
	}
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid data provided",
		})
		return
	}
	transaction.ID = uint(id)
	if err := tc.service.UpdateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Transaction updated successfully",
	})
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description delete a transaction by ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /transactions/{id} [delete]
func (tc *TransactionController) DeleteTransaction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID format",
		})
		return
	}
	if err := tc.service.DeleteTransaction(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transaction",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Transaction deleted successfully",
	})
}
