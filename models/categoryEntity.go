package models

type Category struct {
	ID   uint   `gorm:"column:ID;autoIncrement;primaryKey"`
	Name string `gorm:"column:NAME;type:varchar(255)"`
}

func (Category) TableName() string {
	return "CATEGORY"
}
