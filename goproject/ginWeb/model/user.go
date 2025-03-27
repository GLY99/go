package model

// controller model
type User struct {
	ID    string `json:"id" binding:"required" gorm:"primaryKey"`
	Name  string `json:"name" binding:"required" gorm:"size:255;not null"`
	Age   int    `json:"age" binding:"required,gte=18" gorm:"not nll"`
	Email string `json:"email" binding:"required" gorm:"uniqueIndex;size:255;not null"`
}

// 映射结构体和表的关系
func (user *User) TableName() string {
	return "custom_users"
}
