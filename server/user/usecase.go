package user

import (
	"hackathon-backend/utils/logger"
	"time"

	"firebase.google.com/go/auth"
)

type Usecase interface {
	Register(token *auth.Token, data []byte) error
}

type usecase struct {
	dao Dao
}

func NewUsecase(dao Dao) Usecase {
	return &usecase{
		dao: dao,
	}
}

func (u *usecase) Register(token *auth.Token, data []byte) error {
	userData := UserData{
		UID:            token.UID,
		Username:       token.Claims["name"].(string),
		Email:          token.Claims["email"].(string),
		ProfileContent: nil,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}

	if err := u.dao.Register(userData); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// func (u *usecase) Edit(token *auth.Token, data []byte) error {
// 	userData := UserData{
// 		UID:            token.UID,
// 		Username:       token.Claims["name"].(string),
// 		Email:          token.Claims["email"].(string),
// 		ProfileContent: nil,
// 		CreatedAt:      time.Time{},
// 		UpdatedAt:      time.Time{},
// 	}

// 	if err := u.dao.Edit(userData); err != nil {
// 		logger.Error(err)
// 		return err
// 	}
// 	return nil
// }
