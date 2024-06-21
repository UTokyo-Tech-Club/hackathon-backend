package comment

import (
	"hackathon-backend/utils/logger"

	"firebase.google.com/go/auth"
	"github.com/google/uuid"

	wss "hackathon-backend/server/websocketServer"
)

type Usecase interface {
	Post(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error
}

type usecase struct {
	broadcaster Broadcaster
	dao         Dao
}

func NewUsecase(broadcaster Broadcaster, dao Dao) Usecase {
	return &usecase{
		broadcaster: broadcaster,
		dao:         dao,
	}
}

func (u *usecase) Post(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {

	commentData := CommentData{
		CommentUID:        uuid.New().String(),
		PostUID:           data["postID"].(string),
		Comments:          []string{data["comment"].(string)},
		CommentingUserUID: data["commentingUserUID"].(string),
	}

	if err := u.dao.Post(&commentData); err != nil {
		logger.Error(err)
		return err
	}

	// if err := u.broadcaster.Post(ws, &commentData); err != nil {
	// 	logger.Error(err)
	// 	return err
	// }
	return nil
}
