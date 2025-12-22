package models

import (
	"time"
)

type AdminUser struct{
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	Email string  `json:"email" gorm:"uniqueIndex;size:255"`
	Password string `json:"-" gorm:"column:password_hash;size:255"`
	FullName string  `json:"full_name" gorm:"column:full_name;size:255"`
	Role string  `json:"role" gorm:"size:50"`
	// тоже понять что тут за Tg, id или username
	TelegramID string  `json:"telegram_id" gorm:"column:telegram_id;size:100"`
	IsActive bool `json:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func (AdminUser) TableName() string{
	return "admin_users"
}