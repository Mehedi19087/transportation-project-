package customer

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	 ID uint `json:"id" gorm:"primaryKey"`
	 CustomerName string `json:"customerName" gorm:"not null"`
	 Mobile string `json:"mobile"`
	 Email string `json:"email"`
	 Address string `json:"address"`
	 OpeningBalance string `json:"openingBalance"`
	 Status string `json:"status"`
	 CreatedAt time.Time `json:"created_at"`
	 UpdatedAt time.Time `json:"updated_at"`
     DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}