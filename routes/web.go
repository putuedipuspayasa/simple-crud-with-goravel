package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/repositories/user_repository/user_gorm"
	"goravel/app/services/user_service/user_service_v1"

	"goravel/app/http/controllers"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) {
		ctx.Response().Json(http.StatusOK, http.Json{
			"Hello": "Goravel",
		})
	})

	// user
	userR := user_gorm.NewUserRepository()
	userS := user_service_v1.NewUserService(userR)
	userController := controllers.NewUserController(userS)
	facades.Route().Get("/users", userController.Fetch)
	facades.Route().Get("/user/{uuid}", userController.Show)
	facades.Route().Post("/user", userController.Store)
}
