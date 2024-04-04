package tweet

import (
	"hackathon-backend/mysql"
	"hackathon-backend/neo4j"
	"hackathon-backend/utils/logger"
)

type Dao interface {
	Post(tweet *TweetData) error
	Edit(tweet *TweetData) (*TweetData, error)
	GetNewest(tweet *TweetData, index int) (*TweetData, error)
}

type dao struct{}

func NewDao() Dao {
	return &dao{}
}

func (dao *dao) Post(tweet *TweetData) error {

	// Push tweet data to MySQL
	query := "INSERT INTO tweet (uid, owner_uid, content) VALUES (?, ?, ?)"
	if _, err := mysql.Exec(query, tweet.UID, tweet.OwnerUID, tweet.Content); err != nil {
		logger.Error(err)
		return err
	}

	// Push tweet to Neo4j
	query = "MATCH (u:User {uid: $uid}) CREATE (u)-[:POSTED]->(:Tweet {uid: $tweetUID})"
	if _, err := neo4j.Exec(query, map[string]interface{}{"uid": tweet.OwnerUID, "tweetUID": tweet.UID}); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (dao *dao) Edit(tweet *TweetData) (*TweetData, error) {
	query := "UPDATE tweet SET content = ? WHERE uid = ?"
	if _, err := mysql.Exec(query, tweet.Content, tweet.UID); err != nil {
		logger.Error(err)
		return &TweetData{}, err
	}

	new, err := getTweet(tweet.UID)
	if err != nil {
		logger.Error(err)
		return &TweetData{}, err
	}

	return new, nil
}

func (dao *dao) GetNewest(tweet *TweetData, index int) (*TweetData, error) {

	// Retrieve tweet data
	query := "SELECT * FROM tweet ORDER BY create_time DESC LIMIT 1 OFFSET ?"
	stmt, err := mysql.DB.Prepare(query)
	if err != nil {
		logger.Error(err)
		return tweet, err
	}
	defer stmt.Close()

	rows, err := mysql.DB.Query(query, index)
	if err != nil {
		logger.Error(err)
		return tweet, err
	}
	defer rows.Close()

	rows.Next()
	if err := rows.Scan(&tweet.UID, &tweet.OwnerUID, &tweet.Content, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
		logger.Error(err)
		return tweet, nil
	}

	// Retrieve owner data
	query = "SELECT username, photo_url FROM user WHERE uid = ?"
	stmt, err = mysql.DB.Prepare(query)
	if err != nil {
		logger.Error(err)
		return tweet, err
	}
	defer stmt.Close()

	if err = stmt.QueryRow(tweet.OwnerUID).Scan(&tweet.OwnerUsername, &tweet.OwnerPhotoURL); err != nil {
		logger.Error(err)
		return tweet, err
	}

	// Retrieve number of likes
	query = "MATCH (:User)-[:LIKES]->(t:Tweet {uid: $uid}) RETURN COUNT(t)"
	results, err := neo4j.Exec(query, map[string]interface{}{"uid": tweet.UID})
	if err != nil {
		logger.Error(err)
		return tweet, err
	}
	tweet.NumLikes = int(results[0].Values[0].(int64))
	logger.Info(int(results[0].Values[0].(int64)), "likes")

	return tweet, nil
}

func getTweet(uid string) (*TweetData, error) {
	tweet := &TweetData{}

	// Retrieve tweet data
	query := "SELECT * FROM tweet WHERE uid = ?"
	stmt, err := mysql.DB.Prepare(query)
	if err != nil {
		logger.Error(err)
		return tweet, err
	}
	defer stmt.Close()

	rows, err := mysql.DB.Query(query, uid)
	if err != nil {
		logger.Error(err)
		return tweet, err
	}
	defer rows.Close()

	rows.Next()
	if err := rows.Scan(&tweet.UID, &tweet.OwnerUID, &tweet.Content, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
		logger.Error(err)
		return tweet, nil
	}

	// Retrieve owner data
	query = "SELECT username, photo_url FROM user WHERE uid = ?"
	stmt, err = mysql.DB.Prepare(query)
	if err != nil {
		logger.Error(err)
		return tweet, err
	}
	defer stmt.Close()

	if err = stmt.QueryRow(tweet.OwnerUID).Scan(&tweet.OwnerUsername, &tweet.OwnerPhotoURL); err != nil {
		logger.Error(err)
		return tweet, err
	}

	return tweet, nil
}
