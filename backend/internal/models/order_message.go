package models

import (
	"time"
)

type OrderMessage struct{
	ID int  `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID int  `json:"order_id" gorm:"column:order_id"`
	SenderType string  `json:"sender_type" gorm:"column:sender_type;size:20"`
	Message string `json:"message" gorm:"type:text"`
	IsRead bool `json:"is_read" gorm:"column:is_read;default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`

	//Relations
	Order Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`
}


func (OrderMessage) TableName() string{
	return "order_messages"
}