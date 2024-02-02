package post

type PaginationRequest struct {
	PageId   int `validate:"required"`
	PageSize int `validate:"required"`
}
