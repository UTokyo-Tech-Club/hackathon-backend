package user

type Controller struct {
	usecase Usecase
}

func NewController(usecase Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c *Controller) Register(uid string, email string) (string, error) {
	msg, err := c.usecase.Register(uid, email)
	if err != nil {
		return "", err
	}
	return msg, nil
}
