package auth

import (
	"time"

	"gorm.io/gorm"
)

const (
	RoleAdmin = "admin"
	RoleManager = "manager"
)

type User struct {
	 ID    uint  `json:"id" gorm:"primaryKey"`
	 Name string `json:"name" gorm:"not null"`
	 Password string `json:"-" gorm:"not null"`
	 ProductID uint  `json:"-" gorm:"not null;index"`
	 Role string `json:"role" gorm:"not null;default:'manager'"`
	 Status string `json:"status" gorm:"not null;default:'pending'"`
	 Picture   string    `json:"picture"`
	 CreatedAt time.Time  `json:"created_at"`
	 UpdatedAt time.Time `json:"updated_at"`
	 DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}

// This is what Google will send us
type GoogleUserInfo struct {
    ID      string `json:"id"`      // Google's user ID
    Email   string `json:"email"`   // user@gmail.com
    Name    string `json:"name"`    // "John Doe"
    Picture string `json:"picture"` // Profile photo URL
}

