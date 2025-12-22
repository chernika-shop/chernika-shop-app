package models

import (
	"time"
)

type Wishlist struct{
	ID int  `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID int  `json:"user_id" gorm:"column:user_id"`
	ProductID int  `json:"product_id" gorm:"column:product_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	
	//Relations
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (Wishlist) TableName() string{
	return "wishlist"
}