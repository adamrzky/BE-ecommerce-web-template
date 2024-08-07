package controllers

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/repositories"
	"BE-ecommerce-web-template/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related operations
type UserController struct {
	userService services.UserService
}

// NewUserController returns a new UserController
func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

// GetUserByID retrieves a user by ID
// @Summary Get a user by ID
// @Description Retrieve a user by their ID
// @Tags user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.SuccessResponse{data=models.User} "User retrieved successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /users/{id} [get]
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ctrl.userService.GetUserByID(uint(id))
	if err != nil {
		if err == repositories.ErrUserNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User retrieved successfully",
		Data:    user,
	})
}

// GetUserByUsername retrieves a user by username
// @Summary Get a user by username
// @Description Retrieve a user by their username
// @Tags user
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.SuccessResponse{data=models.User} "User retrieved successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /users/username/{username} [get]
func (ctrl *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := ctrl.userService.GetUserByUsername(username)
	if err != nil {
		if err == repositories.ErrUserNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User retrieved successfully",
		Data:    user,
	})
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided data
// @Tags user
// @Accept json
// @Produce json
// @Param userInput body services.CreateUserInput true "User data"
// @Success 201 {object} models.SuccessResponse{data=models.User} "User created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var input services.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid input",
		})
		return
	}

	user, err := ctrl.userService.CreateUser(input)
	if err != nil {
		if err.Error() == "invalid email address" {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Status:  "error",
				Message: "Invalid email address",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "User created successfully",
		Data:    user,
	})
}

// UpdateUser updates an existing user
// @Summary Update a user
// @Description Update a user with the provided data
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param userInput body services.UpdateUserInput true "User data"
// @Success 200 {object} models.SuccessResponse{data=models.User} "User updated successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input services.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid input",
		})
		return
	}

	user, err := ctrl.userService.UpdateUser(uint(id), input)
	if err != nil {
		if err == repositories.ErrUserNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User updated successfully",
		Data:    user,
	})
}

// DeleteUser deletes a user by ID
// @Summary Delete a user by ID
// @Description Delete a user by their ID
// @Tags user
// @Param id path int true "User ID"
// @Success 204 "User deleted successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.userService.DeleteUser(uint(id))
	if err != nil {
		if err == repositories.ErrUserNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
