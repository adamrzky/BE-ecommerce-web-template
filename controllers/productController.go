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

// post creates a new product
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags Product
// @Accept json
// @Produce json
// @Param product body models.ProductRequest true "Product details"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /product [post]
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

	resp.NewResponseSuccess(c, nil, "data created")
}

// getAll retrieves all products with optional query parameters
// @Summary Get all products
// @Description Retrieve a list of all products with optional filtering by price and pagination
// @Tags Product
// @Accept json
// @Produce json
// @Param min_price query float32 false "Minimum price filter"
// @Param max_price query float32 false "Maximum price filter"
// @Param limit query int true "Limit the number of results returned" mininum(1)
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} models.SuccessResponse{data=[]models.ProductResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /product [get]
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

// getByID retrieves a product by its ID
// @Summary Get a product by ID
// @Description Retrieve a product by its ID
// @Tags Product
// @Produce json
// @Param productID path int true "Product ID"
// @Success 200 {object} models.SuccessResponse{data=models.ProductResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /product/{productID} [get]
func (controller *productController) getByID(c *gin.Context) {
	id := c.Param("productID")

	product, err := controller.productService.GetByID(id)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, product, "data received")
}

// update modifies an existing product
// @Summary Update a product
// @Description Update an existing product with the provided details
// @Tags Product
// @Accept json
// @Produce json
// @Param productID path int true "Product ID"
// @Param product body models.ProductRequest true "Product details"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /product/{productID} [put]
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

	resp.NewResponseSuccess(c, nil, "data updated")
}

// delete removes a product by its ID
// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags Product
// @Param productID path int true "Product ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /product/{productID} [delete]
func (controller *productController) delete(c *gin.Context) {
	id := c.Param("productID")

	if err := controller.productService.Delete(id); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, nil, "data deleted")
}
