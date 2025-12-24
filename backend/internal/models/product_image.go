package models


type ProductImage struct{
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID int `json:"product_id" gorm:"column:product_id"`
	// Вместо полного URL храним только путь в MinIO
	// Например: "products/123/image1.jpg"
	ImagePath string `json:"image_path" gorm:"column:image_path;size:500"`
	SortOrder int `json:"sort_order" gorm:"column:sort_order;default:0"`
	IsMain bool  `json:"is_main" gorm:"column:is_main;default:false"`

	//Relations 
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`

}

func (ProductImage) TableName() string{
	return "product_images"
}

// GetURL возвращает полный URL изображения из MinIO
// Это виртуальное поле, не сохраняется в БД
func (pi *ProductImage) GetURL(minioURL string, bucketName string) string{
	if pi.ImagePath == ""{
		return ""
	}
	return minioURL + "/" + bucketName + "/" + pi.ImagePath
}