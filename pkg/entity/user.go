package entity

import (
	"github.com/ebar-go/egu"
)

type UserEntity struct {
	BaseEntity
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

// TableName 指定模型的表名称
func (UserEntity) TableName() string {
	return TableUser
}

func (entity *UserEntity) ValidatePass(pass string) bool {
	if egu.Md5(pass) == entity.Password {
		return true
	}
	return false
}
