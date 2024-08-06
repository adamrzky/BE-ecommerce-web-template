package controllers

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileService *services.ProfileService
}

func NewProfileController(profileService *services.ProfileService) *ProfileController {
	return &ProfileController{profileService: profileService}
}

// CreateProfile godoc
// @Summary Create New Profile.
// @Description Creating a new Profile.
// @Tags Profile
// @Accept json
// @Produce json
// @Param profileInput body services.ProfileInput true "Profile input"
// @Success 200 {object} models.SuccessResponse{data=models.Profile} "Profile Created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /profiles [post]
func (ctrl *ProfileController) Create(c *gin.Context) {
	var input services.ProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	profile, err := ctrl.profileService.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Profile Created successfully",
		Data:    profile,
	})
}

// UpdateProfile godoc
// @Summary Update Existing Profile.
// @Description Updating an existing Profile by ID.
// @Tags Profile
// @Accept json
// @Produce json
// @Param id path int true "Profile ID"
// @Param profileInput body services.ProfileInput true "Profile input"
// @Success 200 {object} models.SuccessResponse{data=models.Profile} "Profile updated successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 404 {object} models.ErrorResponse "Profile not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /profiles/{id} [put]
func (ctrl *ProfileController) Update(c *gin.Context) {
	profileIDStr := c.Param("id")
	profileID, err := strconv.ParseUint(profileIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid profile ID",
		})
		return
	}

	var input services.ProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	profile, err := ctrl.profileService.Update(uint(profileID), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Profile updated successfully",
		Data:    profile,
	})
}

// GetProfileByID godoc
// @Summary Get Profile by ID.
// @Description Retrieve a Profile by its ID.
// @Tags Profile
// @Accept json
// @Produce json
// @Param id path int true "Profile ID"
// @Success 200 {object} models.SuccessResponse{data=models.Profile} "Profile retrieved successfully"
// @Failure 404 {object} models.ErrorResponse "Profile not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /profiles/{id} [get]
func (ctrl *ProfileController) GetByID(c *gin.Context) {
	profileIDStr := c.Param("id")
	profileID, err := strconv.ParseUint(profileIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Invalid profile ID",
		})
		return
	}

	profile, err := ctrl.profileService.GetProfileByID(uint(profileID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Profile retrieved successfully",
		Data:    profile,
	})
}
