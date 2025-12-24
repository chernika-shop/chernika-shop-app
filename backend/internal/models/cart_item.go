package models

import (
	"time"
)

type CartItem struct{
	ID int  `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID int  `json:"user_id" gorm:"column:user_id"`
	ProductID int  `json:"product_id" gorm:"column:product_id"`
	ProductSizeID int `json:"product_size_id" gorm:"column:product_size_id"`
	Quantity int `json:"quantity" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`

	//Relations
	User User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Product Product  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	ProductSize ProductSize `json:"product_size,omitempty" gorm:"foreignKey:ProductSizeID"`
}

func (CartItem) TableName() string {
	return "cart_items"
}