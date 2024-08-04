package controllers

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register handles user registration
// @Summary Register a new user
// @Description Register a new user with username, password, and email
// @Tags auth
// @Accept json
// @Produce json
// @Param registerInput body services.RegisterInput true "Registration input"
// @Success 200 {object} models.SuccessResponse{data=models.User} "User registered successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var input services.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	user, err := ctrl.authService.Register(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data:    user,
	})
}

// Login handles user login
// @Summary Login a user
// @Description Authenticate a user and return a token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginInput body services.LoginInput true "Login input"
// @Success 200 {object} models.SuccessResponse{data=object{token=string,user=models.User}} "Login successful"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Router /login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var input services.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	token, user, err := ctrl.authService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Login successful",
		Data: map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})
}

// Me retrieves the authenticated user's data
// @Summary Get user data
// @Description Retrieve user data for the currently authenticated user
// @Tags auth
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=models.User} "User data retrieved successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /auth/me [get]
func (ctrl *AuthController) Me(c *gin.Context) {
	userID, err := ctrl.authService.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	user, err := ctrl.authService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User data retrieved successfully",
		Data:    user,
	})
}

// ChangePassword handles changing the user's password
// @Summary Change user password
// @Description Change the password for the currently authenticated user
// @Tags auth
// @Accept json
// @Produce json
// @Param changePasswordInput body services.ChangePasswordInput true "Change password input"
// @Success 200 {object} models.SuccessResponse "Password changed successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /auth/change-password [post]
func (ctrl *AuthController) ChangePassword(c *gin.Context) {
	var input services.ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	userID, err := ctrl.authService.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	err = ctrl.authService.ChangePassword(userID, input.OldPassword, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Password changed successfully",
	})
}
