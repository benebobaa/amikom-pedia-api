package post

import "time"

type ResponsePost struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserId    string    `json:"user_id"`
	RefPostId string    `json:"ref_post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
