package user_request

import "goravel/app/http/requests"

type UserGetRequest struct {
	requests.PaginationRequest
	UUID   *string `json:"uuid" form:"uuid" query:"uuid"`
	Name   *string `json:"name" form:"name" query:"name"`
	Email  *string `json:"email" form:"email" query:"email"`
	Phone  *string `json:"phone" form:"phone" query:"phone"`
	Search *string `json:"search" form:"search" query:"search"`
}
