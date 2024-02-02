package post_controller

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/middleware"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/model/web/post"
	"amikom-pedia-api/module/post/post_service"
	"amikom-pedia-api/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type PostControllerImpl struct {
	PostService post_service.PostService
}

func NewPostController(postService post_service.PostService) PostController {
	return &PostControllerImpl{PostService: postService}
}

func (postController PostControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postCreateRequest := post.RequestPost{}
	authPayload := request.Header.Get(middleware.AuthorizationPayloadKey)
	userNId, _ := utils.FromStringToUsernameAndUUID(authPayload)

	userId := userNId.UserID

	helper.ReadFromRequestBody(request, &postCreateRequest)

	postResponse := postController.PostService.Create(request.Context(), userId, postCreateRequest)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (postController PostControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postUpdateRequest := post.RequestPost{}
	postId := params.ByName("id")

	helper.ReadFromRequestBody(request, &postUpdateRequest)

	authPayload := request.Header.Get(middleware.AuthorizationPayloadKey)
	userNId, _ := utils.FromStringToUsernameAndUUID(authPayload)

	userId := userNId.UserID

	updateResponse := postController.PostService.Update(request.Context(), postId, userId, postUpdateRequest)

	if updateResponse.ID == "" {
		http.Error(writer, "Post not found", http.StatusNotFound)
	}

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (postController PostControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("id")

	postController.PostService.Delete(request.Context(), postId)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (postController PostControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(request.URL.Query().Get("pageSize"))

	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	postResponses := postController.PostService.FindAll(request.Context(), page, pageSize)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponses,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (postController PostControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("id")

	postResponse := postController.PostService.FindById(request.Context(), postId)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}
