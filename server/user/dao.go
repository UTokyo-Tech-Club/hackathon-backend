package user

import (
	"hackathon-backend/mysql"
	"hackathon-backend/neo4j"
	"hackathon-backend/utils/logger"

	"sync"
)

type Dao interface {
	Register(d UserData) error
	Edit(d UserData) error
	GetProfileContent(d *UserData) (*UserData, error)
	PullMetadata(uid string) (*UserData, error)
	Follow(userUID string, targetUID string) error
	Unfollow(userUID string, targetUID string) error
	Bookmark(userUID string, tweetUID string) error
	Unbookmark(userUID string, tweetUID string) error
	Like(userUID string, tweetUID string) error
	Unlike(userUID string, tweetUID string) error
}

type dao struct{}

func NewDao() Dao {
	return &dao{}
}

func (dao *dao) Register(d UserData) error {
	var wg sync.WaitGroup
	wg.Add(2)
	errChan := make(chan error, 2)

	// Push user data to MySQL
	go func() {
		defer wg.Done()

		query := "INSERT IGNORE INTO user (uid, username, email, photo_url) VALUES (?, ?, ?, ?)"
		if _, err := mysql.Exec(query, d.UID, d.Username, d.Email, d.PhotoURL); err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	// Push user node to Neo4j
	go func() {
		defer wg.Done()

		query := "MERGE (:User {uid: $uid})"
		if _, err := neo4j.Exec(query, map[string]interface{}{"uid": d.UID}); err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	return nil
}

func (dao *dao) Edit(d UserData) error {
	query := `UPDATE user 
			SET username = ?, photo_url = ?, profile_content = ?
			WHERE uid = ?`
	if _, err := mysql.Exec(query, d.Username, d.PhotoURL, d.ProfileContent, d.UID); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (dao *dao) GetProfileContent(d *UserData) (*UserData, error) {
	query := "SELECT profile_content FROM user WHERE uid = ?"
	stmt, err := mysql.DB.Prepare(query)
	if err != nil {
		logger.Error(err)
		return d, err
	}
	defer stmt.Close()

	var profileContent []byte
	if err = stmt.QueryRow(d.UID).Scan(&profileContent); err != nil {
		logger.Error(err)
		return d, err
	}

	d.ProfileContent = profileContent
	return d, nil
}

func (dao *dao) PullMetadata(uid string) (*UserData, error) {
	var d UserData
	var followingUsers []string = []string{}
	var bookmarkedTweets []string = []string{}
	var likedTweets []string = []string{}

	// Get following users
	query := `MATCH (:User {uid: $userUID})-[:FOLLOWS]->(u:User)
			RETURN u.uid`
	results, err := neo4j.Exec(query, map[string]interface{}{"userUID": uid})
	if err != nil {
		logger.Error(err)
		return &d, err
	}
	for _, record := range results {
		followingUsers = append(followingUsers, record.Values[0].(string))
	}
	d.FollowingUsers = followingUsers

	// Get bookmarked tweets
	query = `MATCH (:User {uid: $userUID})-[:BOOKMARKS]->(t:Tweet)
			RETURN t.uid`
	results, err = neo4j.Exec(query, map[string]interface{}{"userUID": uid})
	if err != nil {
		logger.Error(err)
		return &d, err
	}
	for _, record := range results {
		bookmarkedTweets = append(bookmarkedTweets, record.Values[0].(string))
	}
	d.BookmarkedTweets = bookmarkedTweets

	// Get liked tweets
	query = `MATCH (:User {uid: $userUID})-[:LIKES]->(t:Tweet)
			RETURN t.uid`
	results, err = neo4j.Exec(query, map[string]interface{}{"userUID": uid})
	if err != nil {
		logger.Error(err)
		return &d, err
	}
	for _, record := range results {
		likedTweets = append(likedTweets, record.Values[0].(string))
	}
	d.LikedTweets = likedTweets

	return &d, nil
}

func (dao *dao) Follow(userUID string, targetUID string) error {
	query := `MATCH (u:User {uid: $userUID}), (t:User {uid: $targetUID})
			MERGE (u)-[:FOLLOWS]->(t)`
	if _, err := neo4j.Exec(query, map[string]interface{}{"userUID": userUID, "targetUID": targetUID}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (dao *dao) Unfollow(userUID string, targetUID string) error {
	query := `MATCH (:User {uid: $userUID})-[f:FOLLOWS]->(:User {uid: $targetUID})
			DELETE f`
	if _, err := neo4j.Exec(query, map[string]interface{}{"userUID": userUID, "targetUID": targetUID}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (dao *dao) Bookmark(userUID string, tweetUID string) error {
	query := `MATCH (u:User {uid: $userUID}), (t:Tweet {uid: $tweetUID})
			MERGE (u)-[:BOOKMARKS]->(t)`
	if _, err := neo4j.Exec(query, map[string]interface{}{"userUID": userUID, "tweetUID": tweetUID}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (dao *dao) Unbookmark(userUID string, tweetUID string) error {
	query := `MATCH (:User {uid: $userUID})-[b:BOOKMARKS]->(:Tweet {uid: $tweetUID})
			DELETE b`
	if _, err := neo4j.Exec(query, map[string]interface{}{"userUID": userUID, "tweetUID": tweetUID}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (dao *dao) Like(userUID string, tweetUID string) error {
	query := `MATCH (u:User {uid: $userUID}), (t:Tweet {uid: $tweetUID})
			MERGE (u)-[:LIKES]->(t)`
	if _, err := neo4j.Exec(query, map[string]interface{}{"userUID": userUID, "tweetUID": tweetUID}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (dao *dao) Unlike(userUID string, tweetUID string) error {
	query := `MATCH (:User {uid: $userUID})-[l:LIKES]->(:Tweet {uid: $tweetUID})
			DELETE l`
	if _, err := neo4j.Exec(query, map[string]interface{}{"userUID": userUID, "tweetUID": tweetUID}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
