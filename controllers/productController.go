package controllers

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/services"
	"BE-ecommerce-web-template/utils/resp"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService services.ProductService
}

func NewProductController(e *gin.Engine, cs services.ProductService) {
	handler := productController{cs}

	productGroup := e.Group("/product")
	{
		productGroup.POST("", handler.post)
		productGroup.GET("", handler.getAll)
		productGroup.GET("/:productID", handler.getByID)
		productGroup.PUT("/:productID", handler.update)
		productGroup.DELETE("/:productID", handler.delete)

	}
}

func (controller *productController) post(c *gin.Context) {
	var req models.ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.NewResponseBadRequest(c, err.Error())
		return
	}

	if err := controller.productService.Post(req); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseWriteSuccess(c, "data created")
}

func (controller *productController) getAll(c *gin.Context) {
	var queryParams models.ProductQueryParam
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	products, err := controller.productService.GetAll(queryParams)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, products, "data received")
}

func (controller *productController) getByID(c *gin.Context) {
	id := c.Param("productID")

	product, err := controller.productService.GetByID(id)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, product, "data received")
}

func (controller *productController) update(c *gin.Context) {
	var req models.ProductRequest
	id := c.Param("productID")

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.NewResponseBadRequest(c, err.Error())
		return
	}

	if err := controller.productService.Update(req, id); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseWriteSuccess(c, "data updated")
}

func (controller *productController) delete(c *gin.Context) {
	id := c.Param("productID")

	if err := controller.productService.Delete(id); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseWriteSuccess(c, "data deleted")
}
