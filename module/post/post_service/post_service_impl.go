package post_service

import (
	"amikom-pedia-api/exception"
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/post"
	"amikom-pedia-api/module/post/post_repository"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository post_repository.PostRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewPostService(postRepository post_repository.PostRepository, DB *sql.DB, validate *validator.Validate) PostService {
	return &PostServiceImpl{PostRepository: postRepository, DB: DB, Validate: validate}
}

func (postService PostServiceImpl) Create(ctx context.Context, id string, requestPost post.RequestPost) post.ResponsePost {
	err := postService.Validate.Struct(requestPost)
	helper.PanicIfError(err)

	tx, err := postService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	requestPostDomain := domain.Post{
		UserId:  id,
		Content: requestPost.Content,
	}

	result := postService.PostRepository.Create(ctx, tx, requestPostDomain)

	return helper.ToPostResponse(result)
}

func (postService PostServiceImpl) Update(ctx context.Context, id string, userId string, requestPost post.RequestPost) post.ResponsePost {
	err := postService.Validate.Struct(requestPost)
	fmt.Println(err)
	helper.PanicIfError(err)

	tx, err := postService.DB.Begin()
	fmt.Println(err)
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	postId, err := postService.PostRepository.FindById(ctx, tx, domain.Post{ID: id})
	fmt.Println(err)
	helper.PanicIfError(err)

	requestPostDomain := domain.Post{
		ID:      postId.ID,
		UserId:  userId,
		Content: requestPost.Content,
	}

	result := postService.PostRepository.Update(ctx, tx, requestPostDomain)
	return helper.ToPostResponse(result)
}

func (postService PostServiceImpl) Delete(ctx context.Context, id string) {
	err := postService.Validate.Var(id, "required")
	helper.PanicIfError(err)

	tx, err := postService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	postId, err := postService.PostRepository.FindById(ctx, tx, domain.Post{ID: id})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	postService.PostRepository.Delete(ctx, tx, postId)
}

func (postService PostServiceImpl) FindAll(ctx context.Context, page int, pageSize int) []post.ResponsePost {
	tx, err := postService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	offset := (page - 1) * pageSize

	arg := post.PaginationParams{
		Limit:  pageSize,
		Offset: offset,
	}

	posts := postService.PostRepository.FindAll(ctx, tx, arg)

	var postResponses []post.ResponsePost
	for _, result := range posts {
		postResponses = append(postResponses, helper.ToPostResponse(result))
	}

	return postResponses
}

func (postService PostServiceImpl) FindById(ctx context.Context, id string) post.ResponsePost {
	err := postService.Validate.Var(id, "required")
	helper.PanicIfError(err)

	tx, err := postService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	postId := domain.Post{ID: id}

	result, err := postService.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPostResponse(result)
}
