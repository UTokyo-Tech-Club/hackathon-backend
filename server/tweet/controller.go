package tweet

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

func (c *Controller) Post(token *auth.Token, data []byte) error {
	if err := c.usecase.Post(token, data); err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Posted tweet: ", token.UID)
	return nil
}
