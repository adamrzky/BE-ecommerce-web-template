package models

import "time"

type Profile struct {
	ID        uint      `gorm:"column:ID;primaryKey"`
	Name      string    `gorm:"column:NAME"`
	Gender    string    `gorm:"column:GENDER"`
	City      string    `gorm:"column:CITY"`
	Date      time.Time `gorm:"column:DATE"`
	Address   string    `gorm:"column:ADDRESS"`
	Phone     string    `gorm:"column:PHONE"`
	UserID    uint      `gorm:"column:user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User User `gorm:"foreignKey:UserID"`
}

type ProfileResponse struct {
	ID        uint               `json:"id"`
	UserID    uint               `json:"user_id"`
	Name      string             `json:"name"`
	Gender    string             `json:"gender"`
	City      string             `json:"city"`
	Date      time.Time          `json:"date"`
	Address   string             `json:"address"`
	Phone     string             `json:"phone"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	User      SimpleUserResponse `json:"user"`
}

func (Profile) TableName() string {
	return "PROFILE"
}
