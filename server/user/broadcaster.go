package user

import (
	wss "hackathon-backend/server/websocketServer"
	"hackathon-backend/utils/logger"
)

type Broadcaster interface {
	Follow(ws *wss.WSS, userUID string, userToFollowUID string) error
}

type broadcaster struct{}

func NewBroadcaster() Broadcaster {
	return &broadcaster{}
}

func (b *broadcaster) Follow(ws *wss.WSS, userUID string, userToFollowUID string) error {

	logger.Info("Processing boradcast follow: ", userUID, " -> ", userToFollowUID)

	client, ok := ws.ClientUIDMap.Load(userToFollowUID)
	if !ok {
		logger.Error("Client not found: ", userUID)
		return nil
	}
	conn := client.(*wss.Client).Conn

	conn.WriteJSON(map[string]interface{}{"type": "user", "action": "follow", "data": map[string]interface{}{"followerUID": userUID}})

	logger.Info("Broadcasted follow: ", userUID, " -> ", userToFollowUID)
	return nil
}
