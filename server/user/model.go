package user

import (
	"encoding/json"
	"time"
)

type UserData struct {
	UID            string          `json:"uid"`
	Username       string          `json:"username"`
	Email          string          `json:"email"`
	ProfileContent json.RawMessage `json:"profileContent"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}
