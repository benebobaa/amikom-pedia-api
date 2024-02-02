package post

type PaginationParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
