package controllers

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/services"
	"BE-ecommerce-web-template/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewController struct {
	service services.ReviewService
}

func NewReviewController(service services.ReviewService) *reviewController {
	return &reviewController{service: service}
}

// GetMyReviews godoc
// @Summary Get all reviews by current authenticated user.
// @Description Get a list of reviews.
// @Tags Review
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=[]models.Review} "Success fetch my reviews"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /my-reviews [get]
func (h *reviewController) GetMyReviews(c *gin.Context) {
	var userID, _ = token.ExtractTokenID(c)

	reviews, err := h.service.GetMyReview(int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Get my reviews success",
		Data:    reviews,
	})	
}

// GetReviewById godoc
// @Summary Get Review.
// @Description Get an Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Success 200 {object} models.SuccessResponse{data=models.Review} "Success fetch review by id"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /reviews/{id} [get]
func (h *reviewController) GetReviewById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	review, err := h.service.GetReviewByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Get review by id success",
		Data:    review,
	})	
}

// GetReviewByProductId godoc
// @Summary Get all reviews by current authenticated user.
// @Description Get a list of reviews.
// @Tags Review
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} models.SuccessResponse{data=[]models.Review} "Success fetch my reviews"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router  /reviews-product/{id} [get]
func (h *reviewController) GetReviewByProductId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	reviews, err := h.service.GetReviewByProductID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Get review by product id success",
		Data:    reviews,
	})	
}

// CreateReview godoc
// @Summary Create New Review.
// @Description Creating a new Review.
// @Tags Review
// @Produce json
// @Param Body body models.ReviewInput true "the body to create a new Review"
// @Success 200 {object} models.SuccessResponse{data=models.Review} "Success create new review"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /reviews [post]
func (h *reviewController) CreateReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var userID, _ = token.ExtractTokenID(c)
	var input models.ReviewInput

	// Validasi Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	// Validasi Foreign Key
    var product models.Product
    if err := db.Where("id = ?", input.ProductID).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Product not found",
		} )
        return
    }

	// Validasi Foreign Key
    var transaction models.Transaction
    if err := db.Where("id = ?", input.TransactionID).First(&transaction).Error; err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		} )
        return
    }

	// Limit user hanya bisa review 1x product
	var review models.Review
    err := db.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&review).Error
	if err != nil && err != gorm.ErrRecordNotFound {
    	// Unexpected error occurred
    	c.JSON(http.StatusInternalServerError, models.ErrorResponse{
       		Status:  "error",
        	Message: "Internal server error",
    	})
    	return
	} else if err == nil {
    	// Review already exists
    	c.JSON(http.StatusBadRequest, models.ErrorResponse{
        	Status:  "error",
        	Message: "User sudah mereview produk berikut",
    	})
    return
	}

	review, err = h.service.CreateReview(input, int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Create Review Success",
		Data:    review,
	})
}

// UpdateReview godoc
// @Summary Update Review.
// @Description Update Review by id (only authenticated user with valid user_id).
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Param Body body models.ReviewInput true "the body to update review"
// @Success 200 {object} models.SuccessResponse{data=models.Review} "Success update review"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /reviews/{id} [put]
func (h *reviewController) UpdateReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var userID, _ = token.ExtractTokenID(c)
	var input models.ReviewInput

	// Validasi Parameter ID And Get model if exist
    var review models.Review
    if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Review not found",
		} )
        return
    }

	// Limit user hanya bisa review 1x product
    err := db.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&review).Error
	if err != nil && err != gorm.ErrRecordNotFound {
    	// Unexpected error occurred
    	c.JSON(http.StatusInternalServerError, models.ErrorResponse{
       		Status:  "error",
        	Message: "Internal server error",
    	})
    	return
	} else if err == nil {
    	// Review already exists
    	c.JSON(http.StatusBadRequest, models.ErrorResponse{
        	Status:  "error",
        	Message: "User sudah mereview produk berikut",
    	})
    return
	}

	// Validasi Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	// Validasi Foreign Key
    var product models.Product
    if err := db.Where("id = ?", input.ProductID).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Product not found",
		} )
        return
    }

	// Validasi Foreign Key
    var transaction models.Transaction
    if err := db.Where("id = ?", input.TransactionID).First(&transaction).Error; err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		} )
        return
    }

	// Check apakah user berhak untuk edit
	if review.UserID != int(userID) {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		} )
        return
	}

	updatedReview, err := h.service.UpdateReview(input, int(userID), int(review.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Update Review Success",
		Data:    updatedReview,
	})
}

// DeleteReview godoc
// @Summary Delete one Review.
// @Description Delete a Review by id (only authenticated user with valid user_id).
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Success 200 {object} models.SuccessResponse "Success delete a review"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /reviews/{id} [delete]
func (h *reviewController) DeleteReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var userID, _ = token.ExtractTokenID(c)
	
	// Validasi Parameter ID And Get model if exist
    var review models.Review
    if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Review not found",
		} )
        return
    }

	// Check apakah user berhak untuk edit
	if review.UserID != int(userID) {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		} )
        return
	}

	err := h.service.DeleteReview(int(review.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Delete Review Success",
		Data:    true,
	})
}