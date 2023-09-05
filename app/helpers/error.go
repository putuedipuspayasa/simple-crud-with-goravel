package helpers

import "github.com/goravel/framework/contracts/http"

func ErrorResponse(ctx http.Context, status int, errorMessage any) {
	ctx.Response().Json(status, http.Json{
		"type":    "error",
		"message": errorMessage,
	})
	return
}
