// route 使用单独的路由模块来管理中间件与路由
package route

import (
	_ "ego-demo/docs"
	"ego-demo/internal/handler"
	"ego-demo/pkg/data"
	egoHandler "github.com/ebar-go/ego/http/handler"
	"github.com/ebar-go/ego/http/middleware"
	"github.com/gin-gonic/gin"
)

// Load
func Load(router *gin.Engine) {
	// 加载中间件,trace,cors,requestLog,recover
	router.Use(middleware.CORS, middleware.RequestLog, middleware.Recover)

	// 通过 {host}/swagger/index.html访问swagger web
	router.GET("/swagger/*any", egoHandler.SwaggerHandler())

	router.GET("/", handler.IndexHandler)

	router.GET("/order/:id", handler.GetOrderHandler)

	// 定义版本，方便版本升级
	v1 := router.Group("v1")
	{
		// 登录
		v1.POST("user/login", handler.UserAuthHandler)

		// 注册
		v1.POST("user/register", handler.UserRegisterHandler)

	}

	// 指定路由组需要JWT校验
	private := v1.Group("private").Use(middleware.JWT(&data.UserClaims{}))
	{
		private.GET("user/info", handler.GetUserInfoHandler)
	}
}
