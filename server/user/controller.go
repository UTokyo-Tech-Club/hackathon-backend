package user

import (
	"hackathon-backend/utils/logger"

	"fmt"

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
		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return err
	}

	msg := `{"error": "null"}`
	ws.WriteMessage(websocket.TextMessage, []byte(msg))

	logger.Info("Registered user: ", token.UID, msg)
	return nil
}

func (c *Controller) Edit(ws *websocket.Conn, token *auth.Token, data []byte) error {
	if err := c.usecase.Edit(token, data); err != nil {
		logger.Error(err)
		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return err
	}

	msg := `{"error": "null"}`
	ws.WriteMessage(websocket.TextMessage, []byte(msg))

	logger.Info("Updated user: ", token.UID, msg)
	return nil
}

func (c *Controller) GetProfileContent(ws *websocket.Conn, token *auth.Token, data []byte) error {
	content, err := c.usecase.GetProfileContent(token, data)
	if err != nil {
		logger.Error(err)
		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"content": "null", "error": "%s"}`, err.Error())))
		return err
	}

	msg := fmt.Sprintf(`{"content": %s, "error": "null"}`, string(content))
	ws.WriteMessage(websocket.TextMessage, []byte(msg))

	logger.Info("Sent profile content: ", token.UID, msg)
	return nil
}
