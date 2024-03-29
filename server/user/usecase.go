package user

import "time"

type Usecase interface {
	Register(uid string, email string) (string, error)
}

type usecase struct {
	dao Dao
}

func NewUsecase(dao Dao) Usecase {
	return &usecase{
		dao: dao,
	}
}

func (u *usecase) Register(uid string, email string) (string, error) {
	userData := UserData{
		UID:            uid,
		Username:       email,
		Email:          email,
		ProfileContent: nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := u.dao.Register(userData); err != nil {
		return "", err
	}
	return "Successfully registered", nil
}
