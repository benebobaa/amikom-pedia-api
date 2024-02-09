package image

import (
	"time"
)

type CreateImageResponse struct {
	ImageType string    `json:"image_type"`
	ImageUrl  string    `json:"image_url"`
	FilePath  string    `json:"file_path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
