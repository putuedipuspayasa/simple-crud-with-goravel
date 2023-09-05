package user_service_v1

import (
	"context"
	"github.com/google/uuid"
	"goravel/app/http/request/user_request"
	"goravel/app/http/responses"
	"goravel/app/models"
	"goravel/app/repositories/user_repository"
)

type userService struct {
	userRepository user_repository.UserRepository
}

func NewUserService(userRepository user_repository.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (s userService) Fetch(ctx context.Context, request user_request.UserGetRequest) (users []models.User, pagination responses.PaginationResponse, err error) {
	users, pagination, err = s.userRepository.Fetch(ctx, request)
	if err != nil {
		return
	}
	return
}

func (s userService) Store(ctx context.Context, request user_request.UserStoreRequest) (user models.User, err error) {
	user.UUID = uuid.New().String()
	user.Name = request.Name
	user.Phone = request.Phone
	user.Email = request.Email
	user.Password = request.Password

	err = s.userRepository.Store(ctx, &user)
	return
}

func (s userService) Get(ctx context.Context, UUID string) (user models.User, err error) {
	user, err = s.userRepository.Get(ctx, UUID)
	if err != nil {
		return
	}
	return
}
