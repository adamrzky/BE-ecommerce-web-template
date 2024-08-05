package resp

import (
	"BE-ecommerce-web-template/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewResponseSuccess(c *gin.Context, result interface{}, message string) {
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    result,
	})
}

// Use NewResponseWriteSuccess for writes method such as POST, UPDATE, DELETE
func NewResponseWriteSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: message,
	})
}

func NewResponseBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, models.SuccessResponse{
		Status:  "bad request",
		Message: message,
	})
}

func NewResponseError(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, models.ErrorResponse{
		Status:  "error",
		Message: err,
	})
}

func NewResponseForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, models.SuccessResponse{
		Status:  "forbidden",
		Message: message,
	})
}

func NewResponseUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, models.SuccessResponse{
		Status:  "unauthorized",
		Message: message,
	})
}
