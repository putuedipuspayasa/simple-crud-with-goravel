package user_gorm

import (
	"context"
	"github.com/goravel/framework/facades"
	"goravel/app/http/request/user_request"
	"goravel/app/http/responses"
	"goravel/app/models"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r userRepository) Fetch(ctx context.Context, request user_request.UserGetRequest) (users []models.User, pagination responses.PaginationResponse, err error) {
	page := 1
	perPage := 10
	if request.Page != nil {
		page = *request.Page
	}

	if request.PerPage != nil {
		perPage = *request.PerPage
	}

	err = facades.Orm().Query().Omit("password").Paginate(page, perPage, &users, &pagination.Total)
	if err != nil {
		return
	}

	pagination.CurrentPage = page
	pagination.PerPage = perPage
	return
}

func (r userRepository) Store(ctx context.Context, user *models.User) (err error) {
	err = facades.Orm().Query().Create(&user)
	return
}

func (r userRepository) Get(ctx context.Context, UUID string) (user models.User, err error) {
	err = facades.Orm().Query().Where("uuid = ?", UUID).First(&user)
	return
}
