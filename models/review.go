package models

import "time"

type Review struct {
	ID            	uint   			`gorm:"primary_key" json:"id"`
	UserID        	int    			`gorm:"column:USER_ID" json:"user_id"`
	ProductID     	int    			`gorm:"column:PRODUCT_ID" json:"product_id"`
	TransactionID 	int    			`gorm:"column:TRX_ID" json:"transaction_id"`
	Content       	string 			`gorm:"column:CONTENT" json:"content"`
	CreatedAt     	time.Time		`json:"created_at"`
	UpdatedAt		time.Time		`json:"updated_at"`
	User          	User        	`json:"-"`
	Product       	Product     	`json:"-"`
	Transaction   	Transaction 	`json:"-"`
}

type ReviewResponse struct {
	ID            	int   					`json:"id"`
	UserID        	int    					`json:"user_id"`
	ProductID     	int    					`json:"product_id"`
	TransactionID 	int    					`json:"transaction_id"`
	Content       	string 					`json:"content"`
	CreatedAt     	time.Time				`json:"created_at"`
	UpdatedAt		time.Time				`json:"updated_at"`
	User          	SimpleUserResponse      `json:"user"`
	Product       	SimpleProductResponse   `json:"product"`
	Transaction   	SimpleTrxResponse 		`json:"transaction"`
}

type ReviewInput struct {
	Content			string 	`json:"content" binding:"required"`
	ProductID 		int		`json:"product_id" binding:"required"`
	TransactionID	int 	`json:"transaction_id" binding:"required"`
}

type SimpleUserResponse struct {
	ID			int 	`json:"id"`
	Username	string 	`json:"username"`
	Email		string 	`json:"email"`
}

type SimpleProductResponse struct {
	ID		int 	`json:"id"`
	Name	string 	`json:"name"`
}

type SimpleTrxResponse struct {
	ID				int 	`json:"id"`
	TransactionID	string	`json:"transaction_id"`
	Status			string	`json:"status"`
}

func (Review) TableName() string {
	return "REVIEW"
}