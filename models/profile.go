package models

import "time"

type Profile struct {
	ID      uint      `gorm:"column:ID;primaryKey"`
	Name    string    `gorm:"column:NAME"`
	Gender  string    `gorm:"column:GENDER"`
	City    string    `gorm:"column:CITY"`
	Date    time.Time `gorm:"column:DATE"`
	Address string    `gorm:"column:ADDRESS"`
	Phone   string    `gorm:"column:PHONE"`
	UserID  uint      `gorm:"column:user_id"`

	User User `gorm:"foreignKey:UserID"`
}

func (Profile) TableName() string {
	return "PROFILE"
}
