package tweet

import (
	"encoding/json"
	"time"
)

type TweetData struct {
	TweetUID  string          `json:"tweetUID"`
	UserUID   string          `json:"userUID"`
	Content   json.RawMessage `json:"content"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"udpatdAt"`
}
