package user_goravel_orm

import (
	"context"
	"github.com/goravel/framework/facades"
	"goravel/app/http/requests/user_request"
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

	query := filter(facades.Orm().Query(), request)
	err = query.Omit("password").Paginate(page, perPage, &users, &pagination.Total)
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

func (r userRepository) Get(ctx context.Context, request user_request.UserGetRequest) (user models.User, err error) {
	query := filter(facades.Orm().Query(), request)
	err = query.FirstOrFail(&user)
	return
}

func (r userRepository) Delete(ctx context.Context, request user_request.UserGetRequest) (err error) {
	query := filter(facades.Orm().Query(), request)
	_, err = query.Delete(&models.User{})
	return
}
