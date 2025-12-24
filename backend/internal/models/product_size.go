package models

type ProductSize struct{
	ID int  `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID int  `json:"product_id" gorm:"column:product_id"`
	SizeName string  `json:"size_name" gorm:"column:size_name;size:50"`
	Quantity int  `json:"quantity" gorm:"default:0"`
	
	//Relations
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (ProductSize) TableName() string{
	return "product_sizes"
}