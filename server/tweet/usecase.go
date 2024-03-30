package tweet

import (
	"hackathon-backend/utils/logger"
	"time"

	"firebase.google.com/go/auth"
	"github.com/google/uuid"
)

type Usecase interface {
	Post(token *auth.Token, data []byte) error
}

type usecase struct {
	dao Dao
}

func NewUsecase(dao Dao) Usecase {
	return &usecase{
		dao: dao,
	}
}

func (u *usecase) Post(token *auth.Token, data []byte) error {

	tweetData := TweetData{
		TweetUID:  uuid.New().String(),
		UserUID:   token.UID,
		Content:   data,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	if err := u.dao.Post(tweetData); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
