package models

type UserProductLikes struct {
	UserID    uint    `gorm:"column:user_id;primaryKey"`
	ProductID uint    `gorm:"column:product_id;primaryKey"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
}

func (UserProductLikes) TableName() string {
	return "user_product_likes"
}
