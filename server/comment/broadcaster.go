package comment

import (
	wss "hackathon-backend/server/websocketServer"
	"hackathon-backend/utils/logger"
)

type Broadcaster interface {
	Post(ws *wss.WSS, comment *CommentData) error
}

type broadcaster struct{}

func NewBroadcaster() Broadcaster {
	return &broadcaster{}
}

func (b *broadcaster) Post(ws *wss.WSS, comment *CommentData) error {

	logger.Info("Processing broadcast comment post: ", comment.CommentUID, " -> all")

	return nil
}
