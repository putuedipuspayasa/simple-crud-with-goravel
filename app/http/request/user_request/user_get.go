package user_request

import "goravel/app/http/request"

type UserGetRequest struct {
	request.PaginationRequest
	UUID  string `json:"uuid" form:"uuid" query:"uuid"`
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
	Phone string `json:"phone" form:"phone" query:"phone"`
}
