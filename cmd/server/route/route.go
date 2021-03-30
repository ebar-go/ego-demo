// route 使用单独的路由模块来管理中间件与路由
package route

import (
	_ "ego-demo/docs"
	"ego-demo/internal/handler"
	"github.com/ebar-go/ego/component/log"
	"github.com/ebar-go/ego/http"
	"github.com/ebar-go/ego/http/middleware"
	"github.com/gin-gonic/gin"
)

// Load
func Loader(router *gin.Engine,
	conf *http.Config,
	logger *log.Logger,
	indexHandler handler.IndexHandler, userHandler handler.UserHandler) {
	// 加载中间件,trace,cors,requestLog,recover
	router.Use(middleware.CORS, middleware.RequestLog(logger), middleware.Recover)

	// 通过 {host}/swagger/index.html访问swagger web
	router.GET("/swagger/*any", middleware.Swagger(conf.Swagger))

	router.GET("/", indexHandler.Index)

	// 定义版本，方便版本升级
	v1 := router.Group("v1")
	{
		// 登录
		v1.POST("user/login", userHandler.Auth)

		// 注册
		v1.POST("user/register", userHandler.Register)

	}
}
