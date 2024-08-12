package services

import (
	"BE-ecommerce-web-template/models"
	repository "BE-ecommerce-web-template/repositories"
	"regexp"
	"strconv"
	"strings"
)

type ProductService interface {
	GetAll(params models.ProductQueryParam) ([]models.ProductResponse, int64, error)
	GetByID(id string) (models.ProductResponse, error)
	Post(req models.ProductRequest) error
	Update(req models.ProductRequest, id string) error
	Delete(id string) error
	PostLike(userID uint, productID string) error
	DeleteLike(userID uint, productID string) error
	GetBySlug(slug string) (models.ProductResponse, error)
	GetLikesByUserID(userID uint) ([]models.UserProductLikesResponse, error)
	CompositeLikeExist(userID uint, productID string) (bool, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(cr repository.ProductRepository) ProductService {
	return &productService{cr}
}

func (service *productService) GetAll(params models.ProductQueryParam) ([]models.ProductResponse, int64, error) {
	products, err := service.productRepo.GetAll(params)
	if err != nil {
		return nil, 0, err
	}

	count, err := service.productRepo.CountProducts(params)
	if err != nil {
		return nil, 0, err
	}

	var productResponses []models.ProductResponse
	for _, product := range products {
		productResponse := models.ProductResponse{
			ID:           product.ID,
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			Slug:         product.Slug,
			ImageURL:     product.ImageURL,
			Status:       product.Status,
			CategoryID:   product.CategoryID,
			Category:     models.CategoryResponse(product.Category),
			ProductProps: models.ProductPropsResponse(product.ProductProps),
		}

		productResponses = append(productResponses, productResponse)
	}

	return productResponses, count, nil
}

func (service *productService) GetByID(id string) (models.ProductResponse, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.ProductResponse{}, err
	}

	product, err := service.productRepo.GetByID(uint(idInt))
	if err != nil {
		return models.ProductResponse{}, err
	}

	productResponse := models.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		Slug:         product.Slug,
		ImageURL:     product.ImageURL,
		Status:       product.Status,
		CategoryID:   product.CategoryID,
		Category:     models.CategoryResponse(product.Category),
		ProductProps: models.ProductPropsResponse(product.ProductProps),
	}

	return productResponse, nil
}

func (service *productService) GetBySlug(slug string) (models.ProductResponse, error) {
	product, err := service.productRepo.GetBySlug(slug)
	if err != nil {
		return models.ProductResponse{}, err
	}

	productResponse := models.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		Slug:         product.Slug,
		ImageURL:     product.ImageURL,
		Status:       product.Status,
		CategoryID:   product.CategoryID,
		Category:     models.CategoryResponse(product.Category),
		ProductProps: models.ProductPropsResponse(product.ProductProps),
	}

	return productResponse, nil
}

func (service *productService) Post(req models.ProductRequest) error {
	slug := strings.ToLower(req.Name)
	slug = strings.ReplaceAll(slug, " ", "-")

	reg := regexp.MustCompile(`[^\w-]+`)
	slug = reg.ReplaceAllString(slug, "")

	newProduct := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
		Slug:        slug,
		Status:      req.Status,
		CategoryID:  req.CategoryID,
	}

	return service.productRepo.Post(newProduct)
}

func (service *productService) Update(req models.ProductRequest, id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	existingProduct, err := service.productRepo.GetByID(uint(idInt))
	if err != nil {
		return err
	}

	if req.Name != "" {
		existingProduct.Name = req.Name

		// update the slug
		slug := strings.ToLower(req.Name)
		slug = strings.ReplaceAll(slug, " ", "-")

		reg := regexp.MustCompile(`[^\w-]+`)
		slug = reg.ReplaceAllString(slug, "")
		existingProduct.Slug = slug
	}
	if req.Description != "" {
		existingProduct.Description = req.Description
	}
	if req.Price != 0 {
		existingProduct.Price = req.Price
	}
	if req.ImageURL != "" {
		existingProduct.ImageURL = req.ImageURL
	}
	if req.Status != 0 {
		existingProduct.Status = req.Status
	}
	if req.CategoryID != 0 {
		existingProduct.CategoryID = req.CategoryID
	}

	return service.productRepo.Update(&existingProduct, uint(idInt))
}

func (service *productService) Delete(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return service.productRepo.Delete(uint(idInt))
}

func (service *productService) PostLike(userID uint, productID string) error {
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return err
	}
	return service.productRepo.PostLike(userID, uint(productIDInt))
}

func (service *productService) DeleteLike(userID uint, productID string) error {
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return err
	}
	return service.productRepo.DeleteLike(userID, uint(productIDInt))
}

func (service *productService) GetLikesByUserID(userID uint) ([]models.UserProductLikesResponse, error) {
	likes, err := service.productRepo.GetLikesByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []models.UserProductLikesResponse
	for _, like := range likes {
		response := models.UserProductLikesResponse{
			UserID:    like.UserID,
			ProductID: like.ProductID,
			Product: models.ProductResponse{
				ID:           like.Product.ID,
				Name:         like.Product.Name,
				Description:  like.Product.Description,
				Price:        like.Product.Price,
				Slug:         like.Product.Slug,
				ImageURL:     like.Product.ImageURL,
				Status:       like.Product.Status,
				CategoryID:   like.Product.CategoryID,
				Category:     models.CategoryResponse(like.Product.Category),
				ProductProps: models.ProductPropsResponse(like.Product.ProductProps),
			},
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (service *productService) CompositeLikeExist(userID uint, productID string) (bool, error) {
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return false, err
	}
	return service.productRepo.CompositeLikeExist(userID, uint(productIDInt))
}
