package dailysidecash

import (
    "time"

    "gorm.io/gorm"
)

type DailySideCash struct {
    ID            uint           `json:"id" gorm:"primaryKey"`
    Date          time.Time      `json:"date" gorm:"type:date;not null;index"` // ✅ Proper date type
    Suzuki        float64        `json:"suzuki"`
    Yamaha        float64        `json:"yamaha"`
    Honda         float64        `json:"honda"`
    HatimRupgonj  float64        `json:"hatim_rupgonj"`
    RakibBenapole float64        `json:"rakib_benapole"`
    MofizBenapole float64        `json:"mofiz_benapole"`
    Aziz          float64        `json:"aziz"`
    Shongram      float64        `json:"shongram"`
    CreatedAt     time.Time      `json:"created_at"`
    UpdatedAt     time.Time      `json:"updated_at"`
    DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

func (DailySideCash) TableName() string {
    return "daily_side_cash"
}