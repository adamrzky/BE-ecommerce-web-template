package services

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/repositories"
)

type RoleService interface {
	GetRoleByID(id uint) (models.Role, error)
	CreateRole(input CreateRoleInput) (models.Role, error)
	UpdateRole(id uint, input UpdateRoleInput) (models.Role, error)
	DeleteRole(id uint) error
}

type roleService struct {
	RoleRepo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) RoleService {
	return &roleService{RoleRepo: repo}
}

type CreateRoleInput struct {
	Name string `json:"name"`
}

type UpdateRoleInput struct {
	Name *string `json:"name,omitempty"`
}

func (s *roleService) GetRoleByID(id uint) (models.Role, error) {
	return s.RoleRepo.GetRoleByID(id)
}

func (s *roleService) CreateRole(input CreateRoleInput) (models.Role, error) {
	role := models.Role{
		Name: input.Name,
	}

	err := s.RoleRepo.CreateRole(role)
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func (s *roleService) UpdateRole(id uint, input UpdateRoleInput) (models.Role, error) {
	role, err := s.RoleRepo.GetRoleByID(id)
	if err != nil {
		return models.Role{}, err
	}

	if input.Name != nil {
		role.Name = *input.Name
	}

	err = s.RoleRepo.UpdateRole(role)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func (s *roleService) DeleteRole(id uint) error {
	return s.RoleRepo.DeleteRole(id)
}
