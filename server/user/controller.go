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

// Register user info to the database
// Will ignore if the user already exists
func (c *Controller) Register(ws *websocket.Conn, token *auth.Token, _ map[string]interface{}) error {
	if err := c.usecase.Register(token); err != nil {
		logger.Error(err)
		ws.WriteJSON(map[string]interface{}{"error": err.Error()})
		return err
	}

	ws.WriteJSON(map[string]interface{}{"error": "null"})

	logger.Info("Registered user: ", token.UID)
	return nil
}

func (c *Controller) Edit(ws *websocket.Conn, token *auth.Token, data map[string]interface{}) error {
	if err := c.usecase.Edit(token, data); err != nil {
		logger.Error(err)
		ws.WriteJSON(map[string]interface{}{"error": err.Error()})
		return err
	}

	ws.WriteJSON(map[string]interface{}{"error": "null"})

	logger.Info("Updated user: ", token.UID)
	return nil
}

func (c *Controller) GetProfileContent(ws *websocket.Conn, token *auth.Token, data map[string]interface{}) error {
	content, err := c.usecase.GetProfileContent(token, data)
	if err != nil {
		logger.Error(err)
		ws.WriteJSON(map[string]interface{}{"content": "{}", "error": err.Error()})
		return err
	}

	ws.WriteJSON(map[string]interface{}{"data": content, "error": "null"})

	logger.Info("Sent profile content: ", token.UID, fmt.Sprintf("%s", content))
	return nil
}

func (c *Controller) Follow(ws *websocket.Conn, token *auth.Token, data map[string]interface{}) error {
	if err := c.usecase.Follow(token, data); err != nil {
		logger.Error(err)
		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return err
	}

	ws.WriteJSON(map[string]interface{}{"error": "null"})

	logger.Info("Followed user: ", token.UID)
	return nil
}
