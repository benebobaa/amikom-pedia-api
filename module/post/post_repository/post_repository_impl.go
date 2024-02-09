package post_repository

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

func (postRepository PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Post {
	SQL := `
		SELECT 
    post.id, 
    post.content, 
    post.user_id, 
    post.ref_post_id, 
    post.created_at AS post_created_at, 
    post.updated_at AS post_updated_at, 
    COALESCE(
        JSONB_AGG(
            JSONB_BUILD_OBJECT(
                'user_id', image.user_uuid, 
                'file_path', image.file_path, 
                'image_type', image.image_type, 
                'image_url', image.image_url, 
                'post_id', image.post_id, 
                'created_at', image.created_at, 
                'updated_at', image.updated_at
            )
        ), 
        '[]'::jsonb
    ) AS images
FROM 
    "post"
LEFT JOIN 
    "image" ON post.id = image.post_id
GROUP BY 
    post.id, post.content, post.user_id, post.ref_post_id, post.created_at, post.updated_at
ORDER BY 
    post.created_at DESC

	`

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var postData domain.Post
		var imagesJSON string
		err := rows.Scan(
			&postData.ID,
			&postData.Content,
			&postData.UserId,
			&postData.RefPostId,
			&postData.CreatedAt,
			&postData.UpdatedAt,
			&imagesJSON,
		)
		helper.PanicIfError(err)

		// Unmarshal the JSON array into the Images field

		err = json.Unmarshal([]byte(imagesJSON), &postData.Images)
		log.Println("Error : ", err)
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
