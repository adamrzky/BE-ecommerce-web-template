package models

type Product struct {
	ID           uint         `gorm:"column:ID;autoIncrement;primaryKey"`
	Name         string       `gorm:"column:NAME;type:varchar(255)"`
	Description  string       `gorm:"column:DESCRIPTION;type:varchar(255)"`
	Price        float32      `gorm:"column:PRICE;type:decimal(10,2)"`
	Slug         string       `gorm:"column:SLUG;type:varchar(255)"`
	ImageURL     string       `gorm:"column:IMAGE_URL;type:varchar(255)"`
	Status       int          `gorm:"column:NAME;type:int"`
	CategoryID   uint         `gorm:"column:CATEGORY_ID;type:int"`
	Category     Category     `gorm:"foreignKey:CategoryID"`
	ProductProps ProductProps `gorm:"foreignKey:ProductID;references:ID"`
}

func (Product) TableName() string {
	return "PRODUCT"
}
