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

	if port == "" {
		logger.Warning("PORT is empty")
		port = "8080"
	}

	wss := websocket.Init()
	wss.SetupRouter()
	wss.SetupEventListeners()

	// Warning:
	// This line must be executed to start the server on GCP
	// Avoid calling fatal logger before this line
	// If MySQL connection fails, the build will fail during deployment
	//
	// Note:
	// Connection is first established with HTTP,
	// then upgraded to WebSocket
	logger.Info("Server is running on port " + port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}
