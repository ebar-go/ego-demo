package repository

import (
	"ego-demo/internal/entity"
	"ego-demo/internal/repository/impl"
	"github.com/ebar-go/ego/component/mysql"
	"go.uber.org/dig"
)

type UserRepo interface {
	// 根据邮箱查找用户
	FindByEmail(email string) (*entity.UserEntity, error)
	// 创建用户
	CreateUser(userEntity *entity.UserEntity) error
}

func Inject(container *dig.Container) {
	_ = container.Provide(func(db mysql.Database) UserRepo{
		return impl.NewUserRepo(db)
	})
}
