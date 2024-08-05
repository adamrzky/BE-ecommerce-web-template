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
	PAY_DATE   time.Time `gorm:"type:datetime"`
	CREATED_AT time.Time `gorm:"type:timestamp"`
	UPDATED_AT time.Time `gorm:"type:timestamp on update current_timestamp"`
}

// DetailTransaksi represents the detail of each transaction
type DetailTransaksi struct {
	ID        uint      `gorm:"primaryKey"`
	TRX_ID    uint      `gorm:"index"`
	ProductID uint      `gorm:"index"`
	QTY       int       `gorm:"type:int"`
	Total     float64   `gorm:"type:decimal(10,2)"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp on update current_timestamp"`
}

// TableName sets the insert table name for this struct type
func (Transaction) TableName() string {
	return "TRANSAKSI"
}

func (DetailTransaksi) TableName() string {
	return "DETAILTRANSAKSI"
}
