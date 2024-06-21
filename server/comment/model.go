package comment

type CommentData struct {
	// usecase - dao
	CommentUID        string   `json:"commentUID"`
	PostUID           string   `json:"postUID"`
	Comments          []string `json:"comments"`
	CommentingUserUID string   `json:"commentingUserUID"`

	Usernames []string `json:"usernames"`
	IconUrls  []string `json:"iconUrls"`
}
