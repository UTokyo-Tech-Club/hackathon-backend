package tweet

type Dao interface {
	Post(tweet TweetData, data []byte) (string, error)
}
