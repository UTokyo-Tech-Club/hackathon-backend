package tweet

import (
	"encoding/json"
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
	var tweetData *TweetData
	if err := json.Unmarshal(data, &tweetData); err != nil {
		return err
	}

	if err := c.usecase.Post(userUID, data); err != nil {
		return err
	}

	return nil
}
