package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/services"
	"BE-ecommerce-web-template/utils/token"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func GetPaymentMethods(c *gin.Context) {
	merchantCode := "DS19954"
	apiKey := "21c42276c6d03ddee20ab69e23deaa10"
	datetime := time.Now().Format("2006-01-02 15:04:05")
	paymentAmount := "10000"
	signatureString := fmt.Sprintf("%s%s%s%s", merchantCode, paymentAmount, datetime, apiKey)
	signatureBytes := sha256.Sum256([]byte(signatureString))
	signature := hex.EncodeToString(signatureBytes[:])

	body := map[string]string{
		"merchantCode": merchantCode,
		"amount":       paymentAmount,
		"datetime":     datetime,
		"signature":    signature,
	}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "https://sandbox.duitku.com/webapi/api/merchant/paymentmethod/getpaymentmethod", bytes.NewReader(bodyBytes))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to request payment methods"})
		return
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(responseBody, &result)

	c.JSON(http.StatusOK, result)
}

// GetAllTransactions godoc
// @Summary Get all transactions.
// @Description Retrieve a list of all transactions
// @Tags transactions
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /transactions/all [get]
func (c *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := c.service.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve all transactions",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Success fetch all transactions",
		Data:    transactions,
	})
}

// GetMyTransactions godoc
// @Summary Get all transactions by current authenticated user.
// @Description Retrieve a list of transactions associated with the authenticated user.
// @Tags transactions
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=[]models.Transaction} "Success fetch my transactions"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /mytransactions [get]
func (c *TransactionController) GetMyTransactions(ctx *gin.Context) {
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "User not authenticated or invalid token",
		})
		return
	}

	transactions, err := c.service.GetMyTransactions(int(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transactions",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Success fetch my transactions",
		Data:    transactions,
	})
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
// @Param transaction body models.TransactionDTO true "Create Transaction"
// @Success 201 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /transactions [post]
func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var req models.TransactionDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error in received data: %+v\n", err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid data provided",
		})
		return
	}

	// Convert DTO to Transaction model
	transaction := models.Transaction{
		TRX_ID:     req.TRX_ID,
		PRODUCT_ID: req.PRODUCT_ID,
		USER_ID:    req.USER_ID,
		STATUS:     req.STATUS,
		TOTAL:      req.TOTAL,
		PAY_DATE:   req.PAY_DATE,
		PAY_TYPE:   req.PAY_TYPE,
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
