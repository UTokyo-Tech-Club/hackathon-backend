package server

import (
	"hackathon-backend/server/tweet"
	"hackathon-backend/server/user"

	"github.com/gorilla/websocket"
)

func NewControllers() (map[string]interface{}, map[string]interface{}) {
	userDao := user.NewDao()
	userBroadcaster := user.NewBroadcaster()
	userUsecase := user.NewUsecase(userBroadcaster, userDao)
	userCtl := user.NewController(userUsecase)

	tweetDao := tweet.NewDao()
	tweetBroadcaster := tweet.NewBroadcaster()
	tweetUsecase := tweet.NewUsecase(tweetBroadcaster, tweetDao)
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
			"get_single": tweetCtl.GetSingle,
		},
	}

	// Actions that require authentication
	// Args: *websocket.Conn, *auth.Token, map[string]interface{}
	ctrlAuth := map[string]interface{}{
		"user": map[string]interface{}{
			"auth":                userCtl.Register,
			"edit":                userCtl.Edit,
			"get_profile_content": userCtl.GetProfileContent,
			"pull_metadata":       userCtl.PullMetadata,
			"follow":              userCtl.Follow,
			"unfollow":            userCtl.Unfollow,
			"bookmark":            userCtl.Bookmark,
			"unbookmark":          userCtl.Unbookmark,
			"like":                userCtl.Like,
			"unlike":              userCtl.Unlike,
		},
		"tweet": map[string]interface{}{
			"post": tweetCtl.Post,
			"edit": tweetCtl.Edit,
		},
	}

	return ctrl, ctrlAuth
}

func pong(ws *websocket.Conn, data map[string]interface{}) error {
	ws.WriteJSON(map[string]interface{}{"data": "pong", "error": "null"})
	return nil
}
