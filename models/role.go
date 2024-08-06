package models

import "time"

type Role struct {
	ID        uint      `gorm:"column:ID;primaryKey"`
	Name      string    `gorm:"column:NAME"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Role) TableName() string {
	return "ROLE"
}
