package models

import (
	"github.com/goravel/framework/database/orm"
)

type Transaction struct {
	orm.Model
	Name   string
	Avatar string
	orm.SoftDeletes
}
