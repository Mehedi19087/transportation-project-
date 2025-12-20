package utils

import (
    "database/sql/driver"
    "fmt"
    "time"
)

// BDTime wraps time.Time and automatically converts to Bangladesh time in JSON
type BDTime struct {
    time.Time
}

// MarshalJSON converts UTC to Bangladesh time when sending JSON response
func (bt BDTime) MarshalJSON() ([]byte, error) {
    if bt.Time.IsZero() {
        return []byte("null"), nil
    }
    
    // Convert to Bangladesh time
    loc, _ := time.LoadLocation("Asia/Dhaka")
    bdTime := bt.Time.In(loc)
    
    // Format as string
    formatted := bdTime.Format("2006-01-02 15:04:05")
    return []byte(`"` + formatted + `"`), nil
}

// UnmarshalJSON parses JSON time (for requests if needed)
func (bt *BDTime) UnmarshalJSON(data []byte) error {
    if string(data) == "null" {
        return nil
    }
    
    str := string(data[1 : len(data)-1]) // Remove quotes
    t, err := time.Parse("2006-01-02 15:04:05", str)
    if err != nil {
        return err
    }
    
    bt.Time = t
    return nil
}

// Scan implements sql.Scanner for database reading
func (bt *BDTime) Scan(value interface{}) error {
    if value == nil {
        bt.Time = time.Time{}
        return nil
    }
    
    if t, ok := value.(time.Time); ok {
        bt.Time = t
        return nil
    }
    
    return fmt.Errorf("cannot scan %T into BDTime", value)
}

// Value implements driver.Valuer for database writing
func (bt BDTime) Value() (driver.Value, error) {
    if bt.Time.IsZero() {
        return nil, nil
    }
    return bt.Time, nil
}