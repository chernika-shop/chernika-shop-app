package models

import (
	"time"
)

type User struct{
	ID	int  `json:"id" gorm:"primaryKey;autoIncrement"`
	Email string `json:"email" gorm:"uniqueIndex;size:255"`
	Phone string `json:"phone" gorm:"uniqueIndex;size:20"`
	Password string `json:"-" gorm:"column:password_hash;size:255"`
	FullName string  `json:"full_name" gorm:"column:full_name;size:255"`
	TelegramID string `json:"telegram_id" gorm:"column:telegram_id;size:100"` // поменять название в случае чего
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string{
	return "users"
}