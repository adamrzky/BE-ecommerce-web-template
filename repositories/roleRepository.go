package repositories

import (
	"BE-ecommerce-web-template/models"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrRoleNotFound = errors.New("role not found")
)

type RoleRepository interface {
	GetRoleByID(id uint) (models.Role, error)
	GetAllRoles() ([]models.Role, error)
	CreateRole(role models.Role) error
	UpdateRole(role models.Role) error
	DeleteRole(id uint) error
	GetOrCreateRoleByName(name string) (models.Role, error)
}

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{DB: db}
}

func (r *roleRepository) GetRoleByID(id uint) (models.Role, error) {
	var role models.Role
	if err := r.DB.First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Role{}, ErrRoleNotFound
		}
		return models.Role{}, err
	}
	return role, nil
}

func (r *roleRepository) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) CreateRole(role models.Role) error {
	return r.DB.Create(&role).Error
}

func (r *roleRepository) UpdateRole(role models.Role) error {
	return r.DB.Save(&role).Error
}

func (r *roleRepository) DeleteRole(id uint) error {
	if err := r.DB.Delete(&models.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) GetOrCreateRoleByName(name string) (models.Role, error) {
	var role models.Role

	err := r.DB.Where("name = ?", name).First(&role).Error
	if err == nil {

		return role, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		role = models.Role{Name: name}
		err = r.DB.Create(&role).Error
		if err != nil {
			return models.Role{}, err
		}

		err = r.DB.Where("name = ?", name).First(&role).Error
		if err != nil {
			return models.Role{}, err
		}
		return role, nil
	}

	return models.Role{}, err
}
