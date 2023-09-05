package user_request

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserStoreRequest struct {
	UUID     string  `json:"-"`
	Name     string  `json:"name" form:"name"`
	Email    string  `json:"email" form:"email"`
	Phone    string  `json:"phone" form:"phone"`
	Password *string `json:"password" form:"password"`
}

func (r *UserStoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserStoreRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":  "required|max_len:150",
		"email": "required|email",
		"phone": "required|number",
	}
}

func (r *UserStoreRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserStoreRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserStoreRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
