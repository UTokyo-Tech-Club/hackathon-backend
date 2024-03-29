package server

import (
	"hackathon-backend/server/tweet"
	"hackathon-backend/server/user"
)

func SetupDatabase() {
	user.CreateTable()
	tweet.CreateTable()
}
