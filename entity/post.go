package entity

import "time"

// Post representation a struct
type Post struct {
	Text    string    `json:"text"`
	UserID  string    `json:"userID"`
	Created time.Time `json:"created"`
}
