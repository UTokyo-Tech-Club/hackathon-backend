package user

import (
	"encoding/json"
	"time"
)

type User struct {
	UID            string          `json:"uid"`
	Username       string          `json:"username"`
	Email          string          `json:"email"`
	ProfileContent json.RawMessage `json:"profileContent"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

type UserAuth struct {
	UID   string `json:"uid"`
	Token string `json:"token"`
}
