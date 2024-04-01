package server

import (
	"hackathon-backend/server/tweet"
	"hackathon-backend/server/user"

	"github.com/gorilla/websocket"
)

func NewControllers() (map[string]interface{}, map[string]interface{}) {
	userDao := user.NewDao()
	userUsecase := user.NewUsecase(userDao)
	userCtl := user.NewController(userUsecase)

	tweetDao := tweet.NewDao()
	tweetUsecase := tweet.NewUsecase(tweetDao)
	tweetCtl := tweet.NewController(tweetUsecase)

	// Actions that require no authentication
	// Args: *websocket.Conn, map[string]interface{}
	ctrl := map[string]interface{}{
		"sys": map[string]interface{}{
			"ping": pong,
		},
		"user": map[string]interface{}{},
		"tweet": map[string]interface{}{
			"get_newest": tweetCtl.GetNewest,
		},
	}

	// Actions that require authentication
	// Args: *websocket.Conn, *auth.Token, map[string]interface{}
	ctrlAuth := map[string]interface{}{
		"user": map[string]interface{}{
			"auth":                userCtl.Register,
			"edit":                userCtl.Edit,
			"get_profile_content": userCtl.GetProfileContent,
		},
		"tweet": map[string]interface{}{
			"post": tweetCtl.Post,
		},
	}

	return ctrl, ctrlAuth
}

func pong(ws *websocket.Conn, data map[string]interface{}) error {
	ws.WriteJSON(map[string]interface{}{"data": "pong", "error": "null"})
	return nil
}
