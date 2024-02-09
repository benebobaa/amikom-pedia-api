package image_service

import (
	"context"
	"mime/multipart"
)

type ImageService interface {
	UploadToS3Profile(ctx context.Context, userID string, imgAvatar, imgHeader *multipart.FileHeader) error
	UploadToS3Post(ctx context.Context, userID, postID string, imgPost *multipart.FileHeader) error
}
