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
	UserID    int       `gorm:"column:USER_ID" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"-"`
}

type ProfileResponse struct {
	ID        int                `json:"id"`
	UserID    int                `json:"user_id"`
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
