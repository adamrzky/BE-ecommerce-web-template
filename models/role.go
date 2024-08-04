package models

type Role struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique;not null"`
	Users []User `gorm:"foreignKey:RoleID"`
}
