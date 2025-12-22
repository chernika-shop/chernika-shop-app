package models

import (
	"time"
	"database/sql"
)

type Product struct{
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string  `json:"name" gorm:"size:255"`
	Slug string  `json:"slug" gorm:"uniqueIndex;size:255"`
	Description sql.NullString `json:"description" gorm:"type:text"`
	Price float64  `json:"price" gorm:"type:decimal(10,2)"`
	OldPrice sql.NullFloat64  `json:"old_price" gorm:"type:decimal(10,2);column:old_price"`
	Articule string `json:"articule" gorm:"uniqueIndex;size:100"`
	CategoryID int  `json:"category_id" gorm:"column:category_id"`
	IsActive bool  `json:"is_active" gorm:"column:is_active;default:true"`
	IsAvailable bool  `json:"is_available" gorm:"column:is_available;default:true"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

	//Relations 
	Category Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Images []ProductImage `json:"images,omitempty" gorm:"foreignKey:ProductID"`
	Sizes []ProductSize `json:"sizes,omitempty" gorm:"foreignKey:ProductID"`

}

func (Product) TableName() string{
	return "products"
}