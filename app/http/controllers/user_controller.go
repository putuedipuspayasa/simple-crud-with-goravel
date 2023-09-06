package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/helpers"
	"goravel/app/http/requests/user_request"
	"goravel/app/services/user_service"
)

type UserController struct {
	//Dependent services
	UserService user_service.UserService
}

func NewUserController(UserService user_service.UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (r *UserController) Fetch(ctx http.Context) {
	fetchUserRequest := user_request.UserGetRequest{}

	err := ctx.Request().Bind(&fetchUserRequest)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	
	users, pagination, err := r.UserService.Fetch(ctx, fetchUserRequest)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.Response().Success().Json(http.Json{
		"data": users,
		"meta": http.Json{
			"current_page": pagination.CurrentPage,
			"per_page":     pagination.PerPage,
			"total_rows":   pagination.Total,
			"last_page":    pagination.Total / int64(pagination.PerPage),
		},
	})
	return
}

func (r *UserController) Store(ctx http.Context) {
	var storeUserRequest user_request.UserStoreRequest
	errors, err := ctx.Request().ValidateRequest(&storeUserRequest)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if errors != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}

	user, err := r.UserService.Store(ctx, storeUserRequest)

	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Response().Success().Json(http.Json{
		"data":    user,
		"message": "success",
	})
	return
}

func (r *UserController) Show(ctx http.Context) {
	uuid := ctx.Request().Route("uuid")
	user, err := r.UserService.Get(ctx, uuid)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Response().Success().Json(http.Json{
		"data":    user,
		"message": "success",
	})
	return
}

func (r *UserController) Delete(ctx http.Context) {
	uuid := ctx.Request().Route("uuid")
	err := r.UserService.Delete(ctx, uuid)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Response().Success().Json(http.Json{
		"data":    nil,
		"message": "success",
	})
	return
}
