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

func (i *ImageRepositoryImpl) FindByUserID(ctx context.Context, tx *sql.Tx, userID string) []domain.Image {

	SQL := `SELECT id, user_uuid, post_id, image_type, image_url, file_path, created_at, updated_at FROM "image" WHERE user_uuid = $1`
	rows, err := tx.QueryContext(ctx, SQL, userID)
	helper.PanicIfError(err)

	defer rows.Close()

	var images []domain.Image
	for rows.Next() {
		var image domain.Image
		err = rows.Scan(&image.ID, &image.UserID, &image.PostID, &image.ImageType, &image.ImageUrl, &image.FilePath, &image.CreatedAt, &image.UpdatedAt)
		helper.PanicIfError(err)
		images = append(images, image)
	}

	return images
}

func (i *ImageRepositoryImpl) FindByPostID(ctx context.Context, tx *sql.Tx, postID string) []domain.Image {

	SQL := `SELECT id, user_uuid, post_id, image_type, image_url, file_path, created_at, updated_at FROM "image" WHERE post_id = $1`
	rows, err := tx.QueryContext(ctx, SQL, postID)
	helper.PanicIfError(err)

	defer rows.Close()

	var images []domain.Image
	for rows.Next() {
		var image domain.Image
		err = rows.Scan(&image.ID, &image.UserID, &image.PostID, &image.ImageType, &image.ImageUrl, &image.FilePath, &image.CreatedAt, &image.UpdatedAt)
		helper.PanicIfError(err)
		images = append(images, image)
	}

	return images
}
