package repositories

import (
	"BE-ecommerce-web-template/models"
	"fmt"

	"gorm.io/gorm"
)

// TransactionRepository defines the interface for transaction database operations
type TransactionRepository interface {
	FindByID(id uint) (*models.Transaction, error)
	GetMyTransactions(userID int) ([]models.Transaction, error)
	Create(transaction *models.Transaction) error
	Update(transaction *models.Transaction) error
	Delete(id uint) error
	GetAllTransactions() ([]models.Transaction, error)
	UpdateTransactionStatus(trxID string, status int) error
}

type transactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository returns a new instance of a transaction repository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Product").Preload("User").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetMyTransactions retrieves all transactions associated with a user ID
func (r *transactionRepository) GetMyTransactions(userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Product").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username", "Email")
	}).Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
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

// UpdateTransactionStatus updates the status of a transaction based on TRX_ID
func (repo *transactionRepository) UpdateTransactionStatus(trxID string, status int) error {
	// Mulai dengan cek keberadaan transaksi terlebih dahulu
	var count int64
	result := repo.db.Model(&models.Transaction{}).Where("TRX_ID = ?", trxID).Count(&count)
	if result.Error != nil {
		return result.Error // Error saat query database
	}
	if count == 0 {
		return fmt.Errorf("tidak ada transaksi dengan TRX_ID: %s", trxID) // Tidak ada transaksi yang ditemukan
	}

	// Jika transaksi ditemukan, lanjutkan dengan update statusnya
	result = repo.db.Model(&models.Transaction{}).Where("TRX_ID = ?", trxID).Update("STATUS", status)
	if result.Error != nil {
		return result.Error // Error saat update
	}

	return nil // Sukses
}
