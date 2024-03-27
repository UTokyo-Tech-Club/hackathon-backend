package main

import (
	"hackathon-backend/redis"
	"hackathon-backend/server/websocket"
	"hackathon-backend/utils/logger"
	"net/http"

	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.NewLogger().Fatal("Error loading .env file", err)
	}
}

func main() {
	logger := logger.NewLogger()

	port := os.Getenv("PORT")
	isPortEmpty := port == ""

	if isPortEmpty {
		logger.Fatal("PORT is empty")
	}

	redis.InitRedis()
	redis.IncrementRedis()

	wss := websocket.Init()
	wss.SetupRoutes()
	wss.SetupEventListeners()
	logger.Info("Server is running on port " + port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}
