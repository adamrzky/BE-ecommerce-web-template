package repositories

import (
	"BE-ecommerce-web-template/models"

	"gorm.io/gorm"
)

// TransactionRepository defines the interface for transaction database operations
type TransactionRepository interface {
	FindByID(id uint) (*models.Transaction, error)
	Create(transaction *models.Transaction) error
	Update(transaction *models.Transaction) error
	Delete(id uint) error
}

type transactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository returns a new instance of a transaction repository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

// FindByID finds a transaction by its id
func (repo *transactionRepository) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	result := repo.db.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Name", "Price")
	}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username", "Email")
	}).First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

// Create adds a new transaction to the database
func (repo *transactionRepository) Create(transaction *models.Transaction) error {
	result := repo.db.Create(transaction)
	return result.Error
}

// Update modifies an existing transaction in the database
func (repo *transactionRepository) Update(transaction *models.Transaction) error {
	result := repo.db.Save(transaction)
	return result.Error
}

// Delete removes a transaction from the database
func (repo *transactionRepository) Delete(id uint) error {
	result := repo.db.Delete(&models.Transaction{}, id)
	return result.Error
}
