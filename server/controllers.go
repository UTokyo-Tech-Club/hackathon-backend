package server

import (
	"hackathon-backend/server/tweet"
	"hackathon-backend/server/user"
)

type Controllers struct {
	User  *user.Controller
	Tweet *tweet.Controller
}

func NewControllers() map[string]interface{} {
	userDao := user.NewDao()
	userUsecase := user.NewUsecase(userDao)
	userCtl := user.NewController(userUsecase)

	tweetDao := tweet.NewDao()
	tweetUsecase := tweet.NewUsecase(tweetDao)
	tweetCtl := tweet.NewController(tweetUsecase)

	return map[string]interface{}{
		"user": map[string]interface{}{
			"auth":                userCtl.Register,
			"edit":                userCtl.Edit,
			"get_profile_content": userCtl.GetProfileContent,
		},
		"tweet": map[string]interface{}{
			"post": tweetCtl.Post,
		},
	}
}
