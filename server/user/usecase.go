package user

import (
	"hackathon-backend/utils/logger"

	wss "hackathon-backend/server/websocketServer"

	"firebase.google.com/go/auth"
)

type Usecase interface {
	Register(token *auth.Token) error
	Edit(token *auth.Token, data map[string]interface{}) error
	GetProfileContent(token *auth.Token, data map[string]interface{}) (map[string]interface{}, error)
	PullMetadata(token *auth.Token) (*UserData, error)
	Follow(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error
	Unfollow(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error
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

func (u *usecase) Register(token *auth.Token) error {

	userData := UserData{
		UID:      token.UID,
		Username: token.Claims["name"].(string),
		Email:    token.Claims["email"].(string),
		PhotoURL: token.Claims["picture"].(string),
	}

	if err := u.dao.Register(userData); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (u *usecase) Edit(token *auth.Token, data map[string]interface{}) error {

	userData := UserData{
		UID:            token.UID,
		Username:       data["username"].(string),
		PhotoURL:       data["photoURL"].(string),
		ProfileContent: []byte(data["profileContent"].(string)),
	}

	if err := u.dao.Edit(userData); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (u *usecase) GetProfileContent(token *auth.Token, _ map[string]interface{}) (map[string]interface{}, error) {

	data := UserData{
		UID: token.UID,
	}

	userData, err := u.dao.GetProfileContent(&data)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return map[string]interface{}{"content": userData.ProfileContent}, nil
}

func (u *usecase) PullMetadata(token *auth.Token) (*UserData, error) {

	userData, err := u.dao.PullMetadata(token.UID)
	if err != nil {
		logger.Error(err)
		return &UserData{}, err
	}

	return userData, nil
}

func (u *usecase) Follow(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {

	if err := u.dao.Follow(token.UID, data["userToFollowUID"].(string)); err != nil {
		logger.Error(err)
		return err
	}

	if err := u.broadcaster.Follow(ws, token.UID, data["userToFollowUID"].(string)); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (u *usecase) Unfollow(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {

	if err := u.dao.Unfollow(token.UID, data["userToUnfollowUID"].(string)); err != nil {
		logger.Error(err)
		return err
	}

	if err := u.broadcaster.Unfollow(ws, token.UID, data["userToUnfollowUID"].(string)); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
