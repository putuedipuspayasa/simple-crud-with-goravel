package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	UUID     string  `json:"uuid"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Password *string `json:"password"`
	orm.SoftDeletes
}
