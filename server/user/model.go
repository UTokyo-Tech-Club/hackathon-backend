package user

import (
	"encoding/json"
	"time"

	firebase "firebase.google.com/go"
)

type UserData struct {
	UID            string          `json:"uid"`
	Username       string          `json:"username"`
	Email          string          `json:"email"`
	ProfileContent json.RawMessage `json:"profileContent"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

type FirebaseUserData struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
}

type UserAuth struct {
	FirebaseApp *firebase.App `json:"firebaseApp"`
	Token       string        `json:"token"`
}
