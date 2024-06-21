package comment

import (
	"hackathon-backend/mysql"
	"hackathon-backend/neo4j"
	"hackathon-backend/utils/logger"
)

type Dao interface {
	Post(comment *CommentData) error
}

type dao struct{}

func NewDao() Dao {
	return &dao{}
}

func (dao *dao) Post(comment *CommentData) error {

	// Push comment data to MySQL
	query := "INSERT INTO comment (uid, post_uid, comment, commenting_user_uid) VALUES (?, ?, ?, ?)"
	if _, err := mysql.Exec(query, comment.CommentUID, comment.PostUID, comment.Comments[0], comment.CommentingUserUID); err != nil {
		logger.Error(err)
		return err
	}

	// Push comment to Neo4j
	query = "MATCH (t:Tweet {uid: $postUID}) CREATE (t)-[:COMMENTED]->(:Comment {uid: $commentUID})"
	if _, err := neo4j.Exec(query, map[string]interface{}{"postUID": comment.PostUID, "commentUID": comment.CommentUID}); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
