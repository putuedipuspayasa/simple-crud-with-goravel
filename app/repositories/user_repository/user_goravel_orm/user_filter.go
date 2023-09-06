package user_goravel_orm

import (
	"github.com/goravel/framework/contracts/database/orm"
	"goravel/app/http/requests/user_request"
	"strings"
)

func filter(query orm.Query, request user_request.UserGetRequest) orm.Query {
	if request.UUID != nil {
		query = query.Where("uuid = ?", *request.UUID)
	}
	if request.Name != nil {
		query = query.Where("name = ?", *request.Name)
	}
	if request.Email != nil {
		query = query.Where("email = ?", *request.Email)
	}
	if request.Phone != nil {
		query = query.Where("phone = ?", *request.Phone)
	}
	if request.Search != nil {
		query = query.Where("(LOWER(name) ILIKE ? OR LOWER(email) ILIKE ? OR phone ILIKE ?)", "%"+strings.ToLower(*request.Search)+"%", "%"+strings.ToLower(*request.Search)+"%", "%"+strings.ToLower(*request.Search)+"%")
	}
	return query
}
