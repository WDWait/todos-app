package model

import (
	"time"
)

// GORM 提供了一个内嵌的 Model，包含了 ID, CreatedAt, UpdatedAt, DeletedAt
// 我们可以直接使用它，或者自己定义。
// 这里我们选择自己定义，以保持与之前的一致性，并精确控制字段名。
type Todo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"  gorm:"not null;size:255"`
	Completed bool      `json:"completed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// 如果需要软删除，可以添加 DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
