package user_repository

import (
	"context"
	"goravel/app/http/requests/user_request"
	"goravel/app/http/responses"
	"goravel/app/models"
)

type UserRepository interface {
	Fetch(ctx context.Context, request user_request.UserGetRequest) (users []models.User, pagination responses.PaginationResponse, err error)
	Get(ctx context.Context, request user_request.UserGetRequest) (user models.User, err error)
	Store(ctx context.Context, user *models.User) (err error)
	Delete(ctx context.Context, request user_request.UserGetRequest) (err error)
}
