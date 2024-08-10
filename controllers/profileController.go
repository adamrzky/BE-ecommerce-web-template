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

type ProfileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) *ProfileController {
	return &ProfileController{profileService: profileService}
}

// CreateProfile godoc
// @Summary Create New Profile.
// @Description Creating a new Profile.
// @Tags Profile
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param profileInput body services.ProfileInput true "Profile input"
// @Success 200 {object} models.SuccessResponse{data=models.Profile} "Profile Created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /profiles [post]
func (ctrl *ProfileController) Create(c *gin.Context) {
	var input services.ProfileInput
	userId, _ := token.ExtractTokenID(c)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	profile, err := ctrl.profileService.Create(int(userId), input)
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
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path int true "Profile ID"
// @Security BearerToken
// @Param profileInput body services.ProfileInput true "Profile input"
// @Success 200 {object} models.SuccessResponse{data=models.Profile} "Profile updated successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 404 {object} models.ErrorResponse "Profile not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /profiles/{id} [put]
func (ctrl *ProfileController) Update(c *gin.Context) {
	profileIDStr := c.Param("id")
	userId, _ := token.ExtractTokenID(c)

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

	profile, err := ctrl.profileService.Update(int(profileID), int(userId), input)
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

// GetProfileById godoc
// @Summary Get Profile.
// @Description Get an Profile by id.
// @Tags Profile
// @Produce json
// @Param id path string true "Profile id"
// @Success 200 {object} models.SuccessResponse{data=models.Profile} "Success fetch Profile by id"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /profiles/{id} [get]
func (ctrl *ProfileController) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	profile, err := ctrl.profileService.GetProfileByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Get profile by id success",
		Data:    profile,
	})
}

// GetMyProfiles godoc
// @Summary Get all profiles by current authenticated user.
// @Description Get a list of profiles.
// @Tags Profile
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=[]models.Profile} "Success fetch my profiles"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /my-profiles [get]
func (h *ProfileController) GetMyProfiles(c *gin.Context) {
	var userID, _ = token.ExtractTokenID(c)

	profiles, err := h.profileService.GetMyProfile(int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Get my profiles success",
		Data:    profiles,
	})
}

// DeleteProfile godoc
// @Summary Delete one Profile.
// @Description Delete a Profile by id (only authenticated user with valid user_id).
// @Tags Profile
// @Produce json
// @Param id path string true "Profile id"
// @Success 200 {object} models.SuccessResponse "Success delete a profile"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /profiles/{id} [delete]
func (h *ProfileController) DeleteProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var userID, _ = token.ExtractTokenID(c)

	// Validasi Parameter ID And Get model if exist
	var profile models.Profile
	if err := db.Where("id = ?", c.Param("id")).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "Profile not found",
		})
		return
	}

	// Check apakah user berhak untuk edit
	if profile.UserID != int(userID) {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	err := h.profileService.DeleteProfile(int(profile.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "Delete Profile Success",
		Data:    true,
	})
}
