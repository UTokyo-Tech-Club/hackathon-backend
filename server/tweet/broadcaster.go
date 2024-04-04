package tweet

import (
	wss "hackathon-backend/server/websocketServer"
	"hackathon-backend/utils/logger"
)

type Broadcaster interface {
	Post(ws *wss.WSS, tweet *TweetData) error
	Edit(ws *wss.WSS, tweet *TweetData) error
}

type broadcaster struct{}

func NewBroadcaster() Broadcaster {
	return &broadcaster{}
}

func (b *broadcaster) Post(ws *wss.WSS, tweet *TweetData) error {

	logger.Info("Processing broadcast tweet post: ", tweet.UID, " -> all")

	ws.Clients.Range(func(key, _ interface{}) bool {
		conn := key.(*wss.Client).Conn
		conn.WriteJSON(map[string]interface{}{"type": "tweet", "action": "post", "data": tweet})
		return true
	})

	return nil
}

func (b *broadcaster) Edit(ws *wss.WSS, tweet *TweetData) error {

	logger.Info("Processing broadcast tweet edit: ", tweet.UID, " -> all")

	ws.Clients.Range(func(key, _ interface{}) bool {
		conn := key.(*wss.Client).Conn
		conn.WriteJSON(map[string]interface{}{"type": "tweet", "action": "edit", "data": tweet})
		return true
	})

	return nil
}
