package user

import (
	"encoding/json"
	"time"
)

type UserData struct {
	UID            string          `json:"uid"`
	Username       string          `json:"username,omitempty"`
	Email          string          `json:"email,omitempty"`
	ProfileContent json.RawMessage `json:"profileContent,omitempty"`
	PhotoURL       string          `json:"photoURL,omitempty"`
	CreatedAt      time.Time       `json:"createdAt,omitempty"`
	UpdatedAt      time.Time       `json:"updatedAt,omitempty"`
}

type UpdateData struct {
	Username       string          `json:"username"`
	PhotoURL       string          `json:"photoURL"`
	ProfileContent json.RawMessage `json:"bio"`
}
