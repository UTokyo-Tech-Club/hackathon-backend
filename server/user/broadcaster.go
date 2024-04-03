package user

import (
	wss "hackathon-backend/server/websocketServer"
	"hackathon-backend/utils/logger"
)

type Broadcaster interface {
	Follow(ws *wss.WSS, userUID string, userToFollowUID string) error
	Unfollow(ws *wss.WSS, userUID string, userToUnfollowUID string) error
}

type broadcaster struct{}

func NewBroadcaster() Broadcaster {
	return &broadcaster{}
}

func (b *broadcaster) Follow(ws *wss.WSS, userUID string, userToFollowUID string) error {

	logger.Info("Processing broadcast follow: ", userUID, " -> ", userToFollowUID)

	client, ok := ws.ClientUIDMap.Load(userToFollowUID)
	if !ok {
		return nil
	}
	conn := client.(*wss.Client).Conn

	conn.WriteJSON(map[string]interface{}{"type": "user", "action": "follow", "data": map[string]interface{}{"followerUID": userUID}})

	logger.Info("Broadcasted follow: ", userUID, " -> ", userToFollowUID)
	return nil
}

func (b *broadcaster) Unfollow(ws *wss.WSS, userUID string, userToUnfollowUID string) error {

	logger.Info("Processing broadcast unfollow: ", userUID, " -> ", userToUnfollowUID)

	client, ok := ws.ClientUIDMap.Load(userToUnfollowUID)
	if !ok {
		return nil
	}
	conn := client.(*wss.Client).Conn

	conn.WriteJSON(map[string]interface{}{"type": "user", "action": "unfollow", "data": map[string]interface{}{"followerUID": userUID}})

	logger.Info("Broadcasted unfollow: ", userUID, " -> ", userToUnfollowUID)
	return nil
}
