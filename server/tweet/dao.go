package tweet

import (
	"hackathon-backend/mysql"
	"hackathon-backend/utils/logger"
)

type Dao interface {
	Post(tweet TweetData) error
}

type dao struct{}

func NewDao() Dao {
	return &dao{}
}

func (dao *dao) Post(tweet TweetData) error {
	query := "INSERT INTO tweet (uid, user_uid, content) VALUES (?, ?, ?)"
	if _, err := mysql.Exec(query, tweet.TweetUID, tweet.UserUID, tweet.Content); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
