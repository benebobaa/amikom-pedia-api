package post

import (
	"amikom-pedia-api/model/web/image"
	"time"
)

type ResponsePost struct {
	ID        string                      `json:"id"`
	Content   string                      `json:"content"`
	UserId    string                      `json:"user_id"`
	RefPostId string                      `json:"ref_post_id"`
	Images    []image.CreateImageResponse `json:"images"`
	CreatedAt time.Time                   `json:"created_at"`
	UpdatedAt time.Time                   `json:"updated_at"`
}
