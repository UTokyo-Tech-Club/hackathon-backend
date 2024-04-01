package tweet

import (
	"hackathon-backend/mysql"
	"hackathon-backend/utils/logger"
)

type Dao interface {
	Post(tweet TweetData) error
	GetNewest(index int) (TweetData, error)
}

type dao struct{}

func NewDao() Dao {
	return &dao{}
}

func (dao *dao) Post(tweet TweetData) error {
	query := "INSERT INTO tweet (uid, owner_uid, content) VALUES (?, ?, ?)"
	if _, err := mysql.Exec(query, tweet.UID, tweet.OwnerUID, tweet.Content); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (dao *dao) GetNewest(index int) (TweetData, error) {
	query := "SELECT * FROM tweet ORDER BY create_time DESC LIMIT 1 OFFSET ?"
	stmt, err := mysql.DB.Prepare(query)
	if err != nil {
		logger.Error(err)
		return TweetData{}, err
	}
	defer stmt.Close()

	rows, err := mysql.DB.Query(query, index)
	if err != nil {
		logger.Error(err)
		return TweetData{}, err
	}

	var tweet TweetData
	// for rows.Next() {
	rows.Next()

	if err := rows.Scan(&tweet.UID, &tweet.OwnerUID, &tweet.Content, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
		logger.Error(err)
		// return TweetData{}, err
		return TweetData{}, nil
	}
	// }
	return tweet, nil
}
