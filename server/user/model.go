package user

import (
	"encoding/json"
	"time"

	firebase "firebase.google.com/go"
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
	FirebaseApp *firebase.App `json:"firebaseApp"`
	Token       string        `json:"token"`
}
