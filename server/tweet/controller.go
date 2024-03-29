package tweet

import (
	"hackathon-backend/utils/logger"
)

type Controller struct {
	usecase Usecase
}

func NewController(usecase Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c *Controller) Post(userUID string, data []byte) error {
	if err := c.usecase.Post(userUID, data); err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Posted tweet: ", userUID)
	return nil
}
