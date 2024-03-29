package user

import (
	"hackathon-backend/utils/logger"
	"time"
)

type Usecase interface {
	Register(uid string, email string) error
}

type usecase struct {
	dao Dao
}

func NewUsecase(dao Dao) Usecase {
	return &usecase{
		dao: dao,
	}
}

func (u *usecase) Register(uid string, email string) error {
	userData := UserData{
		UID:            uid,
		Username:       email,
		Email:          email,
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
