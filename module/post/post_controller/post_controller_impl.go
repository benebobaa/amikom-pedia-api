package post_controller

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/middleware"
	"amikom-pedia-api/model/web"
	"amikom-pedia-api/model/web/post"
	"amikom-pedia-api/module/image/image_service"
	"amikom-pedia-api/module/post/post_service"
	"amikom-pedia-api/utils"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type PostControllerImpl struct {
	PostService  post_service.PostService
	ImageService image_service.ImageService
}

func NewPostController(postService post_service.PostService, imageService image_service.ImageService) PostController {
	return &PostControllerImpl{PostService: postService, ImageService: imageService}
}

func (postController PostControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	authPayload := request.Header.Get(middleware.AuthorizationPayloadKey)
	userNId, _ := utils.FromStringToUsernameAndUUID(authPayload)

	userId := userNId.UserID

	err := request.ParseMultipartForm(10 << 20) // 10 MB limit, adjust as needed
	helper.PanicIfError(err)

	//Form-Data Request
	userUpdateRequest := post.RequestPost{
		Content: request.FormValue("content"),
		// Add other form fields as needed
	}
	_, imgHeaderPost, err := request.FormFile("img_post")
	log.Println("imgHeaderPost", err)

	postResponse := postController.PostService.Create(request.Context(), userId, userUpdateRequest)

	if err == nil && imgHeaderPost != nil {
		postController.ImageService.UploadToS3Post(request.Context(), userId, postResponse.ID, imgHeaderPost)
	}

	baseResponse := web.WebResponse{
		Code:   201,
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
