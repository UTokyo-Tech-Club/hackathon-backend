package server

import (
	"hackathon-backend/server/tweet"
	"hackathon-backend/server/user"
	"hackathon-backend/utils/logger"
)

func SetupDatabase() {
	user.CreateTable()
	tweet.CreateTable()
	logger.Info("Created Tables")
}
