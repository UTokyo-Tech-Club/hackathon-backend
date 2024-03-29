package main

import (
	"hackathon-backend/mysql"
	"hackathon-backend/server"
	"hackathon-backend/server/websocket"
	"hackathon-backend/utils/logger"
	"net/http"

	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Warning("Error loading .env file (ignore if called on GCP)", err)
	}

	mysql.Init()
	server.SetupDatabase()
}

func main() {
	port := os.Getenv("PORT")
	isPortEmpty := port == ""

	if isPortEmpty {
		logger.Fatal("PORT is empty")
	}

	wss := websocket.Init()
	wss.SetupRouter()
	wss.SetupEventListeners()
	logger.Info("Server is running on port " + port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}
