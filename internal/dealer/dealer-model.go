package dealer

import (
	"time"
)
type Dealer struct {
	 ID  uint `json:"id" gorm:"primaryKey"`
	 Name string `json:"dealer_name" gorm:"not null"`
	 Destination string `json:"destination" gorm:"not null"`
	 Status string `json:"status"`
	 CreatedAt time.Time `json:"created_at"`
	 UpdatedAt time.Time `json:"updated_at"`
}