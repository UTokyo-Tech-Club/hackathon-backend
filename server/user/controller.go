package user

import (
	"errors"
	wss "hackathon-backend/server/websocketServer"
	"hackathon-backend/utils/logger"

	"fmt"

	"firebase.google.com/go/auth"
)

type Controller struct {
	usecase Usecase
}

func NewController(usecase Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

// Register user info to the database
// Will ignore if the user already exists
func (c *Controller) Register(ws *wss.WSS, token *auth.Token, _ map[string]interface{}) error {
	client, ok := ws.ClientUIDMap.Load(token.UID)
	if !ok {
		err := errors.New("client not found")
		logger.Error(err)
		return err
	}
	conn := client.(*wss.Client).Conn

	if err := c.usecase.Register(token); err != nil {
		logger.Error(err)
		conn.WriteJSON(map[string]interface{}{"error": err.Error()})
		return err
	}

	conn.WriteJSON(map[string]interface{}{"error": "null"})

	logger.Info("Registered user: ", token.UID)
	return nil
}

func (c *Controller) Edit(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {
	client, ok := ws.ClientUIDMap.Load(token.UID)
	if !ok {
		err := errors.New("client not found")
		logger.Error(err)
		return err
	}
	conn := client.(*wss.Client).Conn

	if err := c.usecase.Edit(token, data); err != nil {
		logger.Error(err)
		conn.WriteJSON(map[string]interface{}{"error": err.Error()})
		return err
	}

	conn.WriteJSON(map[string]interface{}{"error": "null"})

	logger.Info("Updated user: ", token.UID)
	return nil
}

func (c *Controller) GetProfileContent(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {
	client, ok := ws.ClientUIDMap.Load(token.UID)
	if !ok {
		err := errors.New("client not found")
		logger.Error(err)
		return err
	}
	conn := client.(*wss.Client).Conn

	content, err := c.usecase.GetProfileContent(token, data)
	if err != nil {
		logger.Error(err)
		conn.WriteJSON(map[string]interface{}{"content": "{}", "error": err.Error()})
		return err
	}

	conn.WriteJSON(map[string]interface{}{"data": content, "error": "null"})

	logger.Info("Sent profile content: ", token.UID, fmt.Sprintf("%s", content))
	return nil
}

func (c *Controller) PullMetadata(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {
	client, ok := ws.ClientUIDMap.Load(token.UID)
	if !ok {
		err := errors.New("client not found")
		logger.Error(err)
		return err
	}
	conn := client.(*wss.Client).Conn
	userData, err := c.usecase.PullMetadata(token)
	if err != nil {
		logger.Error(err)
		conn.WriteJSON(map[string]interface{}{
			"followingUsers":   userData.FollowingUsers,
			"likedTweets":      userData.LikedTweets,
			"bookmarkedTweets": userData.BookmarkedTweets,
			"error":            err.Error()})
		return err
	}

	conn.WriteJSON(map[string]interface{}{
		"followingUsers":   userData.FollowingUsers,
		"likedTweets":      userData.LikedTweets,
		"bookmarkedTweets": userData.BookmarkedTweets,
		"error":            "null"})

	logger.Info("Pulled metadata for user: ", token.UID)
	return nil
}

func (c *Controller) Follow(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {
	client, ok := ws.ClientUIDMap.Load(token.UID)
	if !ok {
		err := errors.New("client not found")
		logger.Error(err)
		return err
	}
	conn := client.(*wss.Client).Conn

	if err := c.usecase.Follow(ws, token, data); err != nil {
		logger.Error(err)
		conn.WriteJSON(map[string]interface{}{"error": err.Error()})
		return err
	}

	conn.WriteJSON(map[string]interface{}{"error": "null"})

	logger.Info("Followed user: ", token.UID)
	return nil
}

func (c *Controller) Unfollow(ws *wss.WSS, token *auth.Token, data map[string]interface{}) error {
	client, ok := ws.ClientUIDMap.Load(token.UID)
	if !ok {
		err := errors.New("client not found")
		logger.Error(err)
		return err
	}
	conn := client.(*wss.Client).Conn

	if err := c.usecase.Unfollow(ws, token, data); err != nil {
		logger.Error(err)
		conn.WriteJSON(map[string]interface{}{"error": err.Error()})
		return err
	}

	conn.WriteJSON(map[string]interface{}{"error": "null"})

	logger.Info("Unfollowed user: ", token.UID)
	return nil
}
