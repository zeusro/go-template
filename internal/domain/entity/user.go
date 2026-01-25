package entity

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user entity in the domain
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Name     string `json:"name"`
	Active   bool   `gorm:"default:true" json:"active"`
}

// TableName specifies the table name
func (User) TableName() string {
	return "users"
}
