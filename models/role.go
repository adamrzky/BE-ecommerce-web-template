package models

type Role struct {
	ID   uint   `gorm:"column:ID;primaryKey"`
	Name string `gorm:"column:NAME"`
}

func (Role) TableName() string {
	return "ROLE"
}
