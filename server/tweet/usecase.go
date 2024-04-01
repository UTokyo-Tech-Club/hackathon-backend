package tweet

import (
	"hackathon-backend/utils/logger"
	"time"

	"firebase.google.com/go/auth"
	"github.com/google/uuid"
)

type Usecase interface {
	Post(token *auth.Token, data map[string]interface{}) error
	GetNewest(data map[string]interface{}) (*TweetData, error)
}

type usecase struct {
	dao Dao
}

func NewUsecase(dao Dao) Usecase {
	return &usecase{
		dao: dao,
	}
}

func (u *usecase) Post(token *auth.Token, data map[string]interface{}) error {

	tweetData := TweetData{
		UID:       uuid.New().String(),
		OwnerUID:  token.UID,
		Content:   []byte(data["content"].(string)),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	if err := u.dao.Post(tweetData); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (u *usecase) GetNewest(data map[string]interface{}) (*TweetData, error) {
	index := int(data["index"].(float64))

	newTweet := TweetData{}
	tweet, err := u.dao.GetNewest(&newTweet, index)
	if err != nil {
		logger.Error(err)
		return &newTweet, err
	}
	return tweet, nil
}
