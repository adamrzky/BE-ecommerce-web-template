package controllers

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/repositories"
	"BE-ecommerce-web-template/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RoleController handles role-related operations
type RoleController struct {
	roleService services.RoleService
}

// NewRoleController returns a new RoleController
func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

// GetRoleByID retrieves a role by ID
// @Summary Get a role by ID
// @Description Retrieve a role by their ID
// @Tags role
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} models.SuccessResponse{data=models.Role} "Role retrieved successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Role not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /roles/{id} [get]
func (ctrl *RoleController) GetRoleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	role, err := ctrl.roleService.GetRoleByID(uint(id))
	if err != nil {
		if errors.Is(err, repositories.ErrRoleNotFound) {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "Role not found",
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
		Message: "Role retrieved successfully",
		Data:    role,
	})
}

// CreateRole creates a new role
// @Summary Create a new role
// @Description Create a new role with the provided data
// @Tags role
// @Accept json
// @Produce json
// @Param roleInput body services.CreateRoleInput true "Role data"
// @Success 201 {object} models.SuccessResponse{data=models.Role} "Role created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /roles [post]
func (ctrl *RoleController) CreateRole(c *gin.Context) {
	var input services.CreateRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid input",
		})
		return
	}

	role, err := ctrl.roleService.CreateRole(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Status:  "success",
		Message: "Role created successfully",
		Data:    role,
	})
}

// UpdateRole updates an existing role
// @Summary Update a role
// @Description Update a role with the provided data
// @Tags role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param roleInput body services.UpdateRoleInput true "Role data"
// @Success 200 {object} models.SuccessResponse{data=models.Role} "Role updated successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 404 {object} models.ErrorResponse "Role not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /roles/{id} [put]
func (ctrl *RoleController) UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input services.UpdateRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid input",
		})
		return
	}

	role, err := ctrl.roleService.UpdateRole(uint(id), input)
	if err != nil {
		if errors.Is(err, repositories.ErrRoleNotFound) {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "Role not found",
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
		Message: "Role updated successfully",
		Data:    role,
	})
}

// DeleteRole deletes a role by ID
// @Summary Delete a role by ID
// @Description Delete a role by their ID
// @Tags role
// @Param id path int true "Role ID"
// @Success 204 "Role deleted successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Role not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /roles/{id} [delete]
func (ctrl *RoleController) DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.roleService.DeleteRole(uint(id))
	if err != nil {
		if errors.Is(err, repositories.ErrRoleNotFound) {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Status:  "error",
				Message: "Role not found",
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
