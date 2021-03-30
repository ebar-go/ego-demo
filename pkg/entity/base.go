package entity

import "github.com/ebar-go/ego/component/mysql"

type BaseEntity struct {
	mysql.Entity
	IsDeleted int `json:"is_deleted" gorm:"column:is_deleted"`
}

const (
	TableUser = "users"
)

const (
	SoftDeleteCondition = "is_deleted = 0"
)
