package post_repository

import (
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/post"
	"context"
	"database/sql"
)

type PostRepository interface {
	Create(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Delete(ctx context.Context, tx *sql.Tx, post domain.Post)
	FindAll(ctx context.Context, tx *sql.Tx, arg post.PaginationParams) []domain.Post
	FindById(ctx context.Context, tx *sql.Tx, post domain.Post) (domain.Post, error)
}
