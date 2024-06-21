package comment

import (
	"errors"
	"hackathon-backend/utils/logger"

	wss "hackathon-backend/server/websocketServer"

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

func (c *Controller) Post(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {
	client, ok := ws.ClientUIDMap.Load(token.UID)
	if !ok {
		err := errors.New("client not found")
		logger.Error(err)
		return err
	}
	conn := client.(*wss.Client).Conn

	if err := c.usecase.Post(ws, token, data); err != nil {
		logger.Error(err)
		return err
	}

	conn.WriteJSON(map[string]interface{}{"source": data["source"], "error": "null"})

	logger.Info("Posted comment: ", data["postUID"], data["comment"], data["commentingUserUID"])
	return nil
}
