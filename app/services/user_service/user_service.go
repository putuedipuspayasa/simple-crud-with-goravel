package user_service

import (
	"context"
	"goravel/app/http/requests/user_request"
	"goravel/app/http/responses"
	"goravel/app/models"
)

type UserService interface {
	Fetch(ctx context.Context, request user_request.UserGetRequest) (users []models.User, pagination responses.PaginationResponse, err error)
	Store(ctx context.Context, request user_request.UserStoreRequest) (user models.User, err error)
	Get(ctx context.Context, UUID string) (user models.User, err error)
	Delete(ctx context.Context, UUID string) (err error)
}
