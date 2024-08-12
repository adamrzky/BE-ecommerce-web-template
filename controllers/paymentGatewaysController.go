package controllers

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPaymentMethods @Summary Get payment methods
// @Description Retrieves available payment methods from the payment gateway using provided merchant details.
// @Tags Payment
// @Accept json
// @Produce json
// @Param merchantCode body string true "Merchant Code"
// @Param apiKey body string true "API Key"
// @Param amount body string true "Amount to be processed"
// @Success 200 {object} map[string]interface{} "Successfully retrieved payment methods"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /payment-methods [post]
func GetPaymentMethods(c *gin.Context) {
	var requestData struct {
		MerchantCode string `json:"merchantCode"`
		ApiKey       string `json:"apiKey"`
		Amount       string `json:"amount"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	datetime := time.Now().Format("2006-01-02 15:04:05")
	signatureString := fmt.Sprintf("%s%s%s%s", requestData.MerchantCode, requestData.Amount, datetime, requestData.ApiKey)
	signatureBytes := sha256.Sum256([]byte(signatureString))
	signature := hex.EncodeToString(signatureBytes[:])

	body := map[string]string{
		"merchantCode": requestData.MerchantCode,
		"amount":       requestData.Amount,
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

// Inquiry @Summary Perform transaction inquiry
// @Description Performs a transaction inquiry to the payment gateway using transaction details.
// @Tags Payment
// @Accept json
// @Produce json
// @Param merchantCode body string true "Merchant Code for the transaction"
// @Param paymentAmount body int true "Amount of the transaction"
// @Param merchantOrderID body string true "Order ID of the transaction"
// @Param productDetails body string true "Details of the product involved in the transaction"
// @Param email body string true "Customer email"
// @Param paymentMethod body string true "Method of payment"
// @Param apiKey body string true "API Key for authentication"
// @Success 200 {object} map[string]interface{} "Successfully performed inquiry"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /payment-inquiry [post]
func Inquiry(c *gin.Context) {
	var requestData struct {
		MerchantCode    string `json:"merchantCode"`
		PaymentAmount   int    `json:"paymentAmount"`
		MerchantOrderID string `json:"merchantOrderID"`
		ProductDetails  string `json:"productDetails"`
		Email           string `json:"email"`
		PaymentMethod   string `json:"paymentMethod"`
		ApiKey          string `json:"apiKey"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Generate signature
	dataSignature := fmt.Sprintf("%s%s%d%s", requestData.MerchantCode, requestData.MerchantOrderID, requestData.PaymentAmount, requestData.ApiKey)
	signature := md5.Sum([]byte(dataSignature))
	signatureHex := hex.EncodeToString(signature[:])

	body := map[string]interface{}{
		"merchantCode":    requestData.MerchantCode,
		"paymentAmount":   requestData.PaymentAmount,
		"merchantOrderID": requestData.MerchantOrderID,
		"productDetails":  requestData.ProductDetails,
		"email":           requestData.Email,
		"paymentMethod":   requestData.PaymentMethod,
		"signature":       signatureHex,
	}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "https://sandbox.duitku.com/webapi/api/merchant/v2/inquiry", bytes.NewReader(bodyBytes))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send inquiry"})
		return
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(responseBody, &result)

	c.JSON(http.StatusOK, result)
}

// CallbackController handles payment gateway callbacks
// Callback @Summary Receive payment gateway callback
// @Description Receives and logs payment callback data from the payment gateway.
// @Tags Payment
// @Accept json
// @Produce json
// @Param data body object true "Payment Callback Data"
// @Success 200 {object} map[string]interface{} "Successfully received callback"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Router /payment-callback [post]
func PaymentCallback(c *gin.Context) {
	// Read the raw body data
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Log the raw request data
	fmt.Println("Raw request data:", string(bodyBytes))

	// Respond that the request has been received
	c.JSON(http.StatusOK, gin.H{"message": "Request received successfully"})
}
