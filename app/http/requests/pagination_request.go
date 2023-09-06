package requests

type PaginationRequest struct {
	Page    *int `json:"page" form:"page" query:"page"`
	PerPage *int `json:"per_page" form:"per_page" query:"per_page"`
}
