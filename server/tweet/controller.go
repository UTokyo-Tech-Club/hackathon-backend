package tweet

import (
	"hackathon-backend/utils/logger"

	wss "hackathon-backend/server/websocketServer"

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

func (c *Controller) Post(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {
	conn := ws.ClientUIDMap[token.UID].Conn

	if err := c.usecase.Post(token, data); err != nil {
		logger.Error(err)
		return err
	}

	conn.WriteJSON(map[string]interface{}{"error": "null"})

	logger.Info("Posted tweet: ", data["content"])
	return nil
}

func (c *Controller) GetNewest(ws *websocket.Conn, data map[string]interface{}) error {
	tweet, err := c.usecase.GetNewest(data)
	if err != nil {
		logger.Error(err)
		ws.WriteJSON(map[string]interface{}{"data": "{}", "error": err.Error()})
		return err
	}

	ws.WriteJSON(map[string]interface{}{"data": tweet, "error": "null"})

	logger.Info("Sending tweet: ", tweet.OwnerUsername, string(tweet.Content))
	return nil
}
