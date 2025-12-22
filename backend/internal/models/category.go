package models

type Category struct{
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string  `json:"name" gorm:"size:100"`
	Slug string `json:"slug" gorm:"uniqueIndex;size:100"`
	ParentID *int `json:"parent_id" gorm:"column:parent_id"`
	SortOrder int  `json:"sort_order" gorm:"column:sort_order;default:0"`
	IsActive bool `json:"is_active" gorm:"column:is_active;default:true"`
	
	//Relations 
	Parent *Category `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children []Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`
}

func (Category) TableName() string{
	return "categories"
}