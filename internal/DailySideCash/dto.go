package dailysidecash

// CreateDailySideCashDTO is the request body for creating a daily side cash record.
// Expect "YYYY-MM-DD" for Date.
type CreateDailySideCashDTO struct {
	Date          string  `json:"date" binding:"required"`
	Suzuki        float64 `json:"suzuki"`
	Yamaha        float64 `json:"yamaha"`
	Honda         float64 `json:"honda"`
	HatimRupgonj  float64 `json:"hatim_rupgonj"`
	RakibBenapole float64 `json:"rakib_benapole"`
	MofizBenapole float64 `json:"mofiz_benapole"`
	Aziz          float64 `json:"aziz"`
	Shongram      float64 `json:"shongram"`
}

// UpdateDailySideCashDTO is the request body for partial updates.
type UpdateDailySideCashDTO struct {
	Date          *string  `json:"date"`
	Suzuki        *float64 `json:"suzuki"`
	Yamaha        *float64 `json:"yamaha"`
	Honda         *float64 `json:"honda"`
	HatimRupgonj  *float64 `json:"hatim_rupgonj"`
	RakibBenapole *float64 `json:"rakib_benapole"`
	MofizBenapole *float64 `json:"mofiz_benapole"`
	Aziz          *float64 `json:"aziz"`
	Shongram      *float64 `json:"shongram"`
}
