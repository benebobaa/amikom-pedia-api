package image_service

import (
	"context"
	"mime/multipart"
)

type ImageService interface {
	UploadToS3(ctx context.Context, userID string, imgAvatar, imgHeader *multipart.FileHeader) error
}
