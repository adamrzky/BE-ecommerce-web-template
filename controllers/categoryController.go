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

	resp.NewResponseWriteSuccess(c, "data created")
}

func (controller *categoryController) getAll(c *gin.Context) {
	categories, err := controller.categoryService.GetAll()
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, categories, "data received")
}

func (controller *categoryController) getByID(c *gin.Context) {
	id := c.Param("categoryID")

	category, err := controller.categoryService.GetByID(id)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, category, "data received")
}

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

	resp.NewResponseWriteSuccess(c, "data updated")
}

func (controller *categoryController) delete(c *gin.Context) {
	id := c.Param("categoryID")

	if err := controller.categoryService.Delete(id); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseWriteSuccess(c, "data deleted")
}
