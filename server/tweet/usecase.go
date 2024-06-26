package tweet

import (
	"hackathon-backend/utils/logger"
	"time"

	"firebase.google.com/go/auth"
	"github.com/google/uuid"

	wss "hackathon-backend/server/websocketServer"
)

type Usecase interface {
	Post(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error
	Edit(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error
	GetNewest(data map[string]interface{}) (*TweetData, error)
	GetSingle(data map[string]interface{}) (*TweetData, error)
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

	tweetData := TweetData{
		UID:        uuid.New().String(),
		OwnerUID:   token.UID,
		Content:    []byte(data["content"].(string)),
		Link:       data["link"].(string),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		LinksBack:  []string{},
		LinksFront: []string{},
		ImageUrl:   data["imageUrl"].(string),

		OwnerUsername: token.Claims["name"].(string),
		OwnerPhotoURL: token.Claims["picture"].(string),
	}

	if err := u.dao.Post(&tweetData); err != nil {
		logger.Error(err)
		return err
	}

	if err := u.broadcaster.Post(ws, &tweetData); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (u *usecase) Edit(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {

	tweetData := TweetData{
		UID:     data["tweetUID"].(string),
		Content: []byte(data["content"].(string)),
	}

	updated, err := u.dao.Edit(&tweetData)

	if err != nil {
		logger.Error(err)
		return err
	}

	if err := u.broadcaster.Edit(ws, updated); err != nil {
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

func (u *usecase) GetSingle(data map[string]interface{}) (*TweetData, error) {
	uid := data["uid"].(string)

	singleTweet := TweetData{}
	tweet, err := u.dao.GetSingle(&singleTweet, uid)
	if err != nil {
		logger.Error(err)
		return &singleTweet, err
	}
	return tweet, nil
}
