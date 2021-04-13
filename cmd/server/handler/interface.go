package handler

import (
	"ego-demo/cmd/server/handler/impl"
	"ego-demo/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type IndexHandler interface {
	Index(ctx *gin.Context)
}

type UserHandler interface {
	Auth(ctx *gin.Context)
	Register(ctx *gin.Context)
}

func Inject(container *dig.Container) {
	_ = container.Provide(func() IndexHandler{
		return impl.NewIndexHandler()
	})
	_ =  container.Provide(func(us service.UserService) UserHandler{
		return impl.NewUserHandler(us)
	})
}