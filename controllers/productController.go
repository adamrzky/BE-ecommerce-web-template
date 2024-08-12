package controllers

import (
	"BE-ecommerce-web-template/middlewares"
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/services"
	"BE-ecommerce-web-template/utils/resp"
	"BE-ecommerce-web-template/utils/role"
	"BE-ecommerce-web-template/utils/token"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService services.ProductService
}

func NewProductController(e *gin.Engine, cs services.ProductService) {
	handler := productController{cs}

	productGroup := e.Group("/product")
	{
		productGroup.POST("", middlewares.JwtAuthMiddleware(role.Admin, role.Developer), handler.post)
		productGroup.GET("", handler.getAll)
		productGroup.GET("/:productID", handler.getByID)
		productGroup.GET("/slug/:productSlug", handler.getBySlug)
		productGroup.PUT("/:productID", middlewares.JwtAuthMiddleware(role.Admin, role.Developer), handler.update)
		productGroup.DELETE("/:productID", middlewares.JwtAuthMiddleware(role.Admin, role.Developer), handler.delete)
		productGroup.PUT("/:productID/likes", middlewares.JwtAuthMiddleware(), handler.likeProduct)
		productGroup.DELETE("/:productID/likes", middlewares.JwtAuthMiddleware(), handler.dislikeProduct)
		productGroup.GET("/likes", middlewares.JwtAuthMiddleware(), handler.getLikesByUserID)
		productGroup.GET("/:productID/likes/status", middlewares.JwtAuthMiddleware(), handler.getLikeStatus)
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
// @Param productName query string false "Product name filter"
// @Param category query int false "Category filter"
// @Param minPrice query float32 false "Minimum price filter"
// @Param maxPrice query float32 false "Maximum price filter"
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

func (controller *productController) getBySlug(c *gin.Context) {
	slug := c.Param("productSlug")

	product, err := controller.productService.GetBySlug(slug)
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

func (controller *productController) likeProduct(c *gin.Context) {
	productID := c.Param("productID")

	claims, err := token.ExtractClaims(c)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	if err := controller.productService.PostLike(claims.ID, productID); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, nil, "success like the product")
}

func (controller *productController) dislikeProduct(c *gin.Context) {
	productID := c.Param("productID")

	claims, err := token.ExtractClaims(c)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	if err := controller.productService.DeleteLike(claims.ID, productID); err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, nil, "success dislike the product")
}

func (controller *productController) getLikesByUserID(c *gin.Context) {
	claims, err := token.ExtractClaims(c)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	res, err := controller.productService.GetLikesByUserID(claims.ID)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, res, "data received")
}

func (controller *productController) getLikeStatus(c *gin.Context) {
	productID := c.Param("productID")

	claims, err := token.ExtractClaims(c)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	status, err := controller.productService.CompositeLikeExist(claims.ID, productID)
	if err != nil {
		resp.NewResponseError(c, err.Error())
		return
	}

	resp.NewResponseSuccess(c, status, "success received like status")
}
