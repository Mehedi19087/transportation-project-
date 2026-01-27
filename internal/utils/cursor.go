package utils

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

// Cursor represents the data needed to paginate
type Cursor struct {
	LastCreatedAt time.Time `json:"last_created_at"`
	LastID        uint      `json:"last_id"`
}

// EncodeCursor creates a base64 string from the last item's data
func EncodeCursor(lastCreatedAt time.Time, lastID uint) string {
	cursor := Cursor{
		LastCreatedAt: lastCreatedAt,
		LastID:        lastID,
	}
	data, _ := json.Marshal(cursor)
	return base64.StdEncoding.EncodeToString(data)
}

// DecodeCursor parses the base64 cursor string back into a struct
func DecodeCursor(cursorStr string) (*Cursor, error) {
	if cursorStr == "" {
		return nil, nil // First page
	}
	data, err := base64.StdEncoding.DecodeString(cursorStr)
	if err != nil {
		return nil, err
	}
	var cursor Cursor
	err = json.Unmarshal(data, &cursor)
	if err != nil {
		return nil, err
	}
	return &cursor, nil
}


