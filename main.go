package main

import (
	"hackathon-backend/mysql"
	// "hackathon-backend/redis"
	"hackathon-backend/server/websocket"
	"hackathon-backend/utils/logger"
	"net/http"

	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file (ignore if called on GCP)", err)
	}

	mysql.Init()
	// redis.Init()
}

func main() {
	port := os.Getenv("PORT")
	isPortEmpty := port == ""

	if isPortEmpty {
		logger.Fatal("PORT is empty")
	}

	wss := websocket.Init()
	wss.SetupRoutes()
	wss.SetupEventListeners()
	logger.Info("Server is running on port " + port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}
