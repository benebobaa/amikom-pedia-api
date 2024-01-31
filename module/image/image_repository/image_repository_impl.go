package image_repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
)

type ImageRepositoryImpl struct {
}

func NewImageRepository() ImageRepository {
	return &ImageRepositoryImpl{}
}

func (i *ImageRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, image domain.Image) domain.Image {
	SQL := `INSERT INTO "image" (user_uuid, post_id ,image_type, image_url, file_path ) VALUES ($1, $2, $3, $4, $5)`
	_, err := tx.ExecContext(ctx, SQL, image.UserID, image.PostID, image.ImageType, image.ImageUrl, image.FilePath)
	helper.PanicIfError(err)

	return image
}
