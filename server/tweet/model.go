package tweet

import (
	"encoding/json"
	"time"
)

type TweetData struct {
	// usecase - dao
	UID        string          `json:"uid"`
	OwnerUID   string          `json:"ownerUID"`
	Content    json.RawMessage `json:"content"`
	Link       string          `json:"link"`
	LinksBack  []string        `json:"linksBack"`
	LinksFront []string        `json:"linksFront"`
	CreatedAt  time.Time       `json:"createdAt,omitempty"`
	UpdatedAt  time.Time       `json:"udpatdAt,omitempty"`

	// controller - usecase
	OwnerUsername string   `json:"ownerUsername,omitempty"`
	OwnerPhotoURL string   `json:"ownerPhotoURL,omitempty"`
	NumLikes      int      `json:"numLikes,omitempty"`
	NumComments   int      `json:"numComments,omitempty"`
	NumLinks      int      `json:"numLinks,omitempty"`
	NumViews      int      `json:"numViews,omitempty"`
	Links         []string `json:"links,omitempty"`

	Comments                []string `json:"comments,omitempty"`
	CommentingUserUsernames []string `json:"commentingUserUsernames,omitempty"`
	CommentingUserIconUrls  []string `json:"commentingUserIconUrls,omitempty"`
}
