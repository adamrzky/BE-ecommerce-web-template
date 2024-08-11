package models

import "time"

// Transaction represents a financial transaction in the system
type Transaction struct {
	ID         uint      `gorm:"primaryKey"`
	TRX_ID     string    `gorm:"unique"`
	PRODUCT_ID uint      `gorm:"index"`
	USER_ID    uint      `gorm:"index"`
	STATUS     int       `gorm:"type:int"`
	TOTAL      int       `gorm:"type:int"`
	PAY_TYPE   string    `gorm:"type:varchar(255)"`
	PAY_DATE   string    `json:"pay_DATE"`
	CreatedAt  time.Time `json:"CREATED_AT"`
	Product    Product   `gorm:"foreignKey:PRODUCT_ID"`
	User       User      `gorm:"foreignKey:USER_ID"`
	UpdatedAt  time.Time `json:"UPDATED_AT"`
}

// DetailTransaksi represents the detail of each transaction
type DetailTransaksi struct {
	ID        uint    `gorm:"primaryKey"`
	TRX_ID    uint    `gorm:"index"`
	ProductID uint    `gorm:"index"`
	QTY       int     `gorm:"type:int"`
	Total     float64 `gorm:"type:decimal(10,2)"`
}

type TransactionPostRequest struct {
	TRX_ID     string    `json:"trx_ID"`
	PRODUCT_ID uint      `json:"product_ID"`
	USER_ID    uint      `json:"user_ID"`
	STATUS     int       `json:"status"`
	TOTAL      float64   `json:"total"`
	PAY_DATE   time.Time `json:"pay_DATE"`
	PAY_TYPE   string    `json:"pay_TYPE"`
}

// TableName sets the insert table name for this struct type
func (Transaction) TableName() string {
	return "TRANSAKSI"
}

func (DetailTransaksi) TableName() string {
	return "DETAILTRANSAKSI"
}
