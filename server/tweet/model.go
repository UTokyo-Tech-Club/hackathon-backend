package tweet

import "encoding/json"

type TweetData struct {
	TweetUID string          `json:"tweetUID"`
	UserUID  string          `json:"userUID"`
	Content  json.RawMessage `json:"content"`
}
