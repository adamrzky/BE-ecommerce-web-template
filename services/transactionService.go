package services

import (
	"BE-ecommerce-web-template/models"
	"BE-ecommerce-web-template/repositories"
)

// TransactionService defines the interface for business logic concerning transactions
type TransactionService interface {
	GetTransactionByID(id uint) (*models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) error
	UpdateTransaction(transaction *models.Transaction) error
	DeleteTransaction(id uint) error
}

type transactionService struct {
	repo repositories.TransactionRepository
}

// NewTransactionService returns a new instance of a TransactionService
func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo}
}

// GetTransactionByID retrieves a transaction by its ID from the repository
func (s *transactionService) GetTransactionByID(id uint) (*models.Transaction, error) {
	transaction, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// CreateTransaction creates a new transaction
func (s *transactionService) CreateTransaction(transaction *models.Transaction) error {
	// Here, you could add business logic (e.g., validation, enrichment) before saving
	return s.repo.Create(transaction)
}

// UpdateTransaction updates an existing transaction
func (s *transactionService) UpdateTransaction(transaction *models.Transaction) error {
	// Additional business logic can be handled here
	return s.repo.Update(transaction)
}

// DeleteTransaction deletes a transaction by its ID
func (s *transactionService) DeleteTransaction(id uint) error {
	return s.repo.Delete(id)
}
