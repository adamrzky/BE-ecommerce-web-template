package models

import "time"

type User struct {
	ID        uint      `gorm:"column:ID;primaryKey"`
	Username  string    `gorm:"column:USERNAME"`
	Password  string    `gorm:"column:PASSWORD"`
	Email     string    `gorm:"column:EMAIL"`
	RoleID    uint      `gorm:"column:role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Role Role `gorm:"foreignKey:RoleID"`
}

func (User) TableName() string {
	return "USER"
}
