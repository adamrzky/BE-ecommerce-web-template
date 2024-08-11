package models

type ProductProps struct {
	ID            uint    `gorm:"column:id;autoIncrement;primaryKey"`
	TotalLikes    uint    `gorm:"column:total_likes"`
	TotalReviews  uint    `gorm:"column:total_reviews"`
	AverageRating float32 `gorm:"column:average_rating"`
	ProductID     uint    `gorm:"column:product_id"`
}

func (ProductProps) TableName() string {
	return "product_props"
}
