package user

import (
	"hackathon-backend/utils/logger"

	"firebase.google.com/go/auth"
)

type Controller struct {
	usecase Usecase
}

func NewController(usecase Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c *Controller) Register(token *auth.Token, data []byte) error {
	if err := c.usecase.Register(token, data); err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Registered user: ", token.UID)
	return nil
}
