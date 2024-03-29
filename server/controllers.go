package server

import (
	"hackathon-backend/server/tweet"
	"hackathon-backend/server/user"
)

type Controllers struct {
	User  *user.Controller
	Tweet *tweet.Controller
}

func NewControllers() *Controllers {
	userDao := user.NewDao()
	userUsecase := user.NewUsecase(userDao)
	userController := user.NewController(userUsecase)

	tweetDao := tweet.NewDao()
	tweetUsecase := tweet.NewUsecase(tweetDao)
	tweetController := tweet.NewController(tweetUsecase)

	return &Controllers{
		User:  userController,
		Tweet: tweetController,
	}
}
