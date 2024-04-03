package tweet

import (
	wss "hackathon-backend/server/websocketServer"
	"hackathon-backend/utils/logger"
)

type Broadcaster interface {
	Post(ws *wss.WSS, tweet TweetData) error
}

type broadcaster struct{}

func NewBroadcaster() Broadcaster {
	return &broadcaster{}
}

func (b *broadcaster) Post(ws *wss.WSS, tweet TweetData) error {

	logger.Info("Processing broadcast tweet post: ", tweet.UID, " -> all")

	ws.Clients.Range(func(key, _ interface{}) bool {
		conn := key.(*wss.Client).Conn
		if key.(*wss.Client).UID != tweet.OwnerUID {
			conn.WriteJSON(map[string]interface{}{"type": "tweet", "action": "post", "data": tweet})
		}
		return true
	})

	return nil
}
