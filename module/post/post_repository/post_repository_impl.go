package post_repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/post"
	"context"
	"database/sql"
	"fmt"
)

type PostRepositoryImpl struct {
}

func NewPostRepository() PostRepository {
	return &PostRepositoryImpl{}
}

func (p PostRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := `INSERT INTO "post" (content, user_id, ref_post_id) VALUES ($1, $2, $3) RETURNING id, content, user_id, ref_post_id, created_at, updated_at`
	row := tx.QueryRowContext(ctx, SQL, post.Content, post.UserId, post.RefPostId)

	err := row.Scan(&post.ID, &post.Content, &post.UserId, &post.RefPostId, &post.CreatedAt, &post.UpdatedAt)
	helper.PanicIfError(err)

	return post
}

func (postRepository PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := `UPDATE "post" SET content = $1 , updated_at = CURRENT_TIMESTAMP WHERE id = $2 AND user_id = $3`
	_, err := tx.ExecContext(ctx, SQL, post.Content, post.ID, post.UserId)
	fmt.Println("Erorr : ", err)
	helper.PanicIfError(err)

	return post
}

func (postRepository PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, post domain.Post) {
	SQL := `DELETE FROM "post" WHERE id = $1`
	_, err := tx.ExecContext(ctx, SQL, post.ID)
	helper.PanicIfError(err)
}

func (postRepository PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, arg post.PaginationParams) []domain.Post {
	SQL := `SELECT id, content, user_id, ref_post_id, created_at, updated_at FROM "post" ORDER BY "created_at" DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, SQL, arg.Limit, arg.Offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		postData := domain.Post{}
		err := rows.Scan(&postData.ID, &postData.Content, &postData.UserId, &postData.RefPostId, &postData.CreatedAt, &postData.UpdatedAt)
		helper.PanicIfError(err)
		posts = append(posts, postData)
	}

	return posts
}

func (postRepository PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, post domain.Post) (domain.Post, error) {
	SQL := `SELECT id, content, user_id, ref_post_id, created_at, updated_at FROM "post" WHERE id = $1`
	rows, err := tx.QueryContext(ctx, SQL, post.ID)
	helper.PanicIfError(err)
	defer rows.Close()

	postData := domain.Post{}
	if rows.Next() {
		rows.Scan(&postData.ID, &postData.Content, &postData.UserId, &postData.RefPostId, &postData.CreatedAt, &postData.UpdatedAt)
		return postData, nil
	} else {
		return postData, err
	}
}
