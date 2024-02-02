package post

type RequestPost struct {
	Content string `json:"content" validate:"required"`
}
