package user

import (
	"encoding/json"
	"time"
)

// Used to interact between usecase and dao
type UserData struct {
	UID            string          `json:"uid"`
	Username       string          `json:"username,omitempty"`
	Email          string          `json:"email,omitempty"`
	ProfileContent json.RawMessage `json:"profileContent,omitempty"`
	PhotoURL       string          `json:"photoURL,omitempty"`
	CreatedAt      time.Time       `json:"createdAt,omitempty"`
	UpdatedAt      time.Time       `json:"updatedAt,omitempty"`
}
