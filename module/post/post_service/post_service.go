package post_service

import (
	"amikom-pedia-api/model/web/post"
	"context"
)

type PostService interface {
	Create(ctx context.Context, id string, requestPost post.RequestPost) post.ResponsePost
	Update(ctx context.Context, id string, userId string, requestPost post.RequestPost) post.ResponsePost
	Delete(ctx context.Context, id string)
	FindAll(ctx context.Context, page int, pageSize int) []post.ResponsePost
	FindById(ctx context.Context, id string) post.ResponsePost
}
