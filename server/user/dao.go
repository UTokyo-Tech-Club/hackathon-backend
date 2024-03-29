package user

import (
	"hackathon-backend/mysql"
	"hackathon-backend/utils/logger"
)

type Dao interface {
	Register(d UserData) error
}

type dao struct{}

func NewDao() Dao {
	return &dao{}
}

func (dao *dao) Register(d UserData) error {
	query := "INSERT IGNORE INTO user (uid, username, email) VALUES (?, ?, ?)"
	if _, err := mysql.Exec(query, d.UID, d.Username, d.Email); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
