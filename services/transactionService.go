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
	GetMyTransactions(userID int) ([]models.Transaction, error)
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
	return s.repo.Create(transaction)
}

// UpdateTransaction updates an existing transaction
func (s *transactionService) UpdateTransaction(transaction *models.Transaction) error {
	return s.repo.Update(transaction)
}

// DeleteTransaction deletes a transaction by its ID
func (s *transactionService) DeleteTransaction(id uint) error {
	return s.repo.Delete(id)
}

// GetMyTransactions retrieves all transactions associated with a user ID
func (s *transactionService) GetMyTransactions(userID int) ([]models.Transaction, error) {
	transactions, err := s.repo.GetMyTransactions(userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
