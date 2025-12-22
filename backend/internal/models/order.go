package models

import (
	"time"
	"database/sql"
)

type Order struct{
	ID int  `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderNumber string  `json:"order_number" gorm:"uniqueIndex;size:50;column:order_number"`
	UserID int  `json:"user_id" gorm:"column:user_id"`
	Status string  `json:"status" gorm:"size:50"`
	TotalAmount float64  `json:"total_amount" gorm:"type:decimal(10,2);column:total_amount"`
	CustomerName string `json:"customer_name" gorm:"column:customer_name;size:255"`
	CustomerPhone string  `json:"customer_phone" gorm:"column:customer_phone;size:20"`
	CustomerEmail string `json:"customer_email" gorm:"column:customer_email;size:255"`
	CustomerTelegram sql.NullString `json:"customer_telegram" gorm:"column:customer_telegram;size:100"`
	PickupPoint sql.NullString `json:"pickup_point" gorm:"column:pickup_point;size:500"`
	Notes sql.NullString `json:"notes" gorm:"type:text"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`

	//Relations
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Items []OrderItem  `json:"items,omitempty" gorm:"foreignKey:OrderID"`
	Messages []OrderMessage  `json:"messages,omitempty" gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string{
	return "orders"
}