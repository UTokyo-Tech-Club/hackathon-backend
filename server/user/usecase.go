package user

import (
	"encoding/json"
	"hackathon-backend/utils/logger"

	"firebase.google.com/go/auth"
)

type Usecase interface {
	Register(token *auth.Token, data []byte) error
	Edit(token *auth.Token, data []byte) error
	GetProfileContent(token *auth.Token, data []byte) ([]byte, error)
}

type usecase struct {
	dao Dao
}

func NewUsecase(dao Dao) Usecase {
	return &usecase{
		dao: dao,
	}
}

func (u *usecase) Register(token *auth.Token, _ []byte) error {

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

func (u *usecase) Edit(token *auth.Token, data []byte) error {

	var newData *UserData
	if err := json.Unmarshal(data, &newData); err != nil {
		logger.Error(err)
		return err
	}

	userData := UserData{
		UID:            token.UID,
		Username:       newData.Username,
		PhotoURL:       newData.PhotoURL,
		ProfileContent: newData.ProfileContent,
	}

	if err := u.dao.Edit(userData); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (u *usecase) GetProfileContent(token *auth.Token, _ []byte) ([]byte, error) {

	data := UserData{
		UID: token.UID,
	}

	userData, err := u.dao.GetProfileContent(data)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return userData.ProfileContent, nil
}
