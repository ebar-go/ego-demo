package repository

import (
	"ego-demo/pkg/data"
	"ego-demo/pkg/entity"
)

type UserRepo interface {
	FindByEmail(email string) (*entity.UserEntity, error)
	BuildUserData(userEntity *entity.UserEntity) data.User
}

type TokenRepo interface {
	CreateToken(user data.User) (string, error)
	ValidateToken(token string) (data.User, error)
}
