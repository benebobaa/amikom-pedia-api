package image_repository

import (
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
)

type ImageRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.Image) domain.Image
}
