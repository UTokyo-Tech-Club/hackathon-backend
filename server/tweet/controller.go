package tweet

import (
	"errors"
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

	logger.Info("Posted tweet: ", data["content"], data["link"])
	return nil
}

func (c *Controller) Edit(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {
	client, ok := ws.ClientUIDMap.Load(token.UID)
	if !ok {
		err := errors.New("client not found")
		logger.Error(err)
		return err
	}
	conn := client.(*wss.Client).Conn

	if err := c.usecase.Edit(ws, token, data); err != nil {
		logger.Error(err)
		return err
	}

	conn.WriteJSON(map[string]interface{}{"source": data["source"], "error": "null"})

	logger.Info("Edited tweet: ", data["tweetUID"])
	return nil
}

func (c *Controller) GetNewest(ws *websocket.Conn, data map[string]interface{}) error {
	tweet, err := c.usecase.GetNewest(data)
	if err != nil {
		logger.Error(err)
		ws.WriteJSON(map[string]interface{}{"source": data["source"], "data": "{}", "error": err.Error()})
		return err
	}

	ws.WriteJSON(map[string]interface{}{"source": data["source"], "data": tweet, "error": "null"})

	logger.Info("Sending tweet: ", tweet.OwnerUsername, string(tweet.Content), tweet.LinksFront, tweet.LinksBack)
	return nil
}

func (c *Controller) GetSingle(ws *websocket.Conn, data map[string]interface{}) error {
	tweet, err := c.usecase.GetSingle(data)
	if err != nil {
		logger.Error(err)
		ws.WriteJSON(map[string]interface{}{"source": data["source"], "data": "{}", "error": err.Error()})
		return err
	}

	ws.WriteJSON(map[string]interface{}{"source": data["source"], "data": tweet, "error": "null"})

	logger.Info("Sending tweet: ", tweet.OwnerUsername, string(tweet.Content), tweet.LinksFront, tweet.LinksBack)
	return nil
}
