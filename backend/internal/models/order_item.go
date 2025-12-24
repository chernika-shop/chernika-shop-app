package models

type OrderItem struct{
	ID int  `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID int  `json:"order_id" gorm:"column:order_id"`
	ProductID int  `json:"product_id" gorm:"column:product_id"`
	ProductSizeID int  `json:"product_size_id" gorm:"column:product_size_id"`
	ProductName string `json:"product_name" gorm:"column:product_name;size:255"`
	ProductPrice float64 `json:"product_price" gorm:"type:decimal(10,2);column:product_price"`
	Quantity int  `json:"quantity"`
	TotalPrice float64 `json:"total_price" gorm:"type:decimal(10,2);column:total_price"`

	//Relations
	Order Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	ProductSize ProductSize `json:"product_size,omitempty" gorm:"foreignKey:ProductSizeID"`
}

func (OrderItem) TableName() string{
	return "order_items"
}