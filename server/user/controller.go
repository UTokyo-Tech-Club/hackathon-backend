package user

import (
	"hackathon-backend/utils/logger"

	"firebase.google.com/go/auth"
	"github.com/gorilla/websocket"
)

type Controller struct {
	usecase Usecase
}

func NewController(usecase Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c *Controller) Register(ws *websocket.Conn, token *auth.Token, data []byte) error {
	if err := c.usecase.Register(token, data); err != nil {
		logger.Error(err)
		return err
	}

	ws.WriteMessage(websocket.TextMessage, []byte(`{"error": "null"}`))

	logger.Info("Registered user: ", token.UID)
	return nil
}

// func (c *Controller) Edit(token *auth.Token, data []byte) error {
// 	if err := c.usecase.Edit(token, data); err != nil {
// 		logger.Error(err)
// 		return err
// 	}

// 	logger.Info("Edited user: ", token.UID)
// 	return nil
// }
