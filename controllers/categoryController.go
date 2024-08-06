package controllers

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/services"
	"BE-ecommerce-web-template/utils/resp"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(e *gin.Engine, cs services.CategoryService) {
	handler := categoryController{cs}

	categoryGroup := e.Group("/category")
	{
		categoryGroup.POST("", handler.post)
		categoryGroup.GET("", handler.getAll)
		categoryGroup.GET("/:categoryID", handler.getByID)
		categoryGroup.PUT("/:categoryID", handler.update)
		categoryGroup.DELETE("/:categoryID", handler.delete)

	}
}

// post creates a new category
// @Summary Create a new category
// @Description Create a new category with the provided details
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.CategoryRequest true "Category details"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /category [post]
func (controller *categoryController) post(c *gin.Context) {
	var req models.CategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.NewResponseBadRequest(c, err.Error())
		return
	}

	if err := controller.categoryService.Post(req); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, nil, "data created")
}

// getAll retrieves all categories
// @Summary Get all categories
// @Description Retrieve a list of all categories
// @Tags Category
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=[]models.CategoryResponse}
// @Failure 500 {object} models.ErrorResponse
// @Router /category [get]
func (controller *categoryController) getAll(c *gin.Context) {
	categories, err := controller.categoryService.GetAll()
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, categories, "data received")
}

// getByID retrieves a category by its ID
// @Summary Get a category by ID
// @Description Retrieve a category by its ID
// @Tags Category
// @Produce json
// @Param categoryID path int true "Category ID"
// @Success 200 {object} models.SuccessResponse{data=models.CategoryResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /category/{categoryID} [get]
func (controller *categoryController) getByID(c *gin.Context) {
	id := c.Param("categoryID")

	category, err := controller.categoryService.GetByID(id)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, category, "data received")
}

// update modifies an existing category
// @Summary Update a category
// @Description Update an existing category with the provided details
// @Tags Category
// @Accept json
// @Produce json
// @Param categoryID path int true "Category ID"
// @Param category body models.CategoryRequest true "Category details"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /category/{categoryID} [put]
func (controller *categoryController) update(c *gin.Context) {
	var req models.CategoryRequest
	id := c.Param("categoryID")

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.NewResponseBadRequest(c, err.Error())
		return
	}

	if err := controller.categoryService.Update(req, id); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, nil, "data updated")
}

// delete removes a category by its ID
// @Summary Delete a category
// @Description Delete a category by its ID
// @Tags Category
// @Param categoryID path int true "Category ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /category/{categoryID} [delete]
func (controller *categoryController) delete(c *gin.Context) {
	id := c.Param("categoryID")

	if err := controller.categoryService.Delete(id); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, nil, "data deleted")
}
