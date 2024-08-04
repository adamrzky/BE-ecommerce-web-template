package models

type User struct {
	ID       uint   `gorm:"column:ID;primaryKey"`
	Username string `gorm:"column:USERNAME"`
	Password string `gorm:"column:PASSWORD"`
	Email    string `gorm:"column:EMAIL"`
	RoleID   uint   `gorm:"column:role_id"`

	Role Role `gorm:"foreignKey:RoleID"`
}

func (User) TableName() string {
	return "user"
}
