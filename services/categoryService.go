package services

import (
	"BE-ecommerce-web-template/models"
	repository "BE-ecommerce-web-template/repositories"
	"strconv"
)

type CategoryService interface {
	GetAll() ([]models.CategoryResponse, error)
	GetByID(id string) (models.CategoryResponse, error)
	Post(req models.CategoryRequest) error
	Update(req models.CategoryRequest, id string) error
	Delete(id string) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(cr repository.CategoryRepository) CategoryService {
	return &categoryService{cr}
}

func (service *categoryService) GetAll() ([]models.CategoryResponse, error) {
	categories, err := service.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var categoryResponses []models.CategoryResponse
	for _, category := range categories {
		categoryResponse := models.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		}

		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return categoryResponses, nil
}

func (service *categoryService) GetByID(id string) (models.CategoryResponse, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.CategoryResponse{}, err
	}

	category, err := service.categoryRepo.GetByID(uint(idInt))
	if err != nil {
		return models.CategoryResponse{}, err
	}

	categoryResponse := models.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return categoryResponse, nil
}

func (service *categoryService) Post(req models.CategoryRequest) error {
	newCategory := models.Category{
		Name: req.Name,
	}

	return service.categoryRepo.Post(newCategory)
}

func (service *categoryService) Update(req models.CategoryRequest, id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	existingCategory, err := service.categoryRepo.GetByID(uint(idInt))
	if err != nil {
		return err
	}

	if req.Name != "" {
		existingCategory.Name = req.Name
	}

	return service.categoryRepo.Update(&existingCategory, uint(idInt))
}

func (service *categoryService) Delete(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return service.categoryRepo.Delete(uint(idInt))
}
