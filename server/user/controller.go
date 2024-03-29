package user

import "hackathon-backend/utils/logger"

type Controller struct {
	usecase Usecase
}

func NewController(usecase Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c *Controller) Register(uid string, email string) error {
	if err := c.usecase.Register(uid, email); err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Registered user: ", uid)
	return nil
}
