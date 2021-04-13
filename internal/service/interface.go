package service

import (
	"ego-demo/internal/dto/request"
	"ego-demo/internal/dto/response"
	"ego-demo/internal/repository"
	"ego-demo/internal/service/impl"
	"github.com/ebar-go/ego/component/auth"
	"go.uber.org/dig"
)

type UserService interface {
	// 校验
	Auth(req request.UserAuthRequest) (*response.UserAuthResponse, error)
	// 注册
	Register(req request.UserRegisterRequest) error
}

func Inject(container *dig.Container) {
	_ =  container.Provide(func(repo repository.UserRepo, jwt auth.Jwt) UserService  {
		return impl.NewUserService(repo, jwt)
	})
}