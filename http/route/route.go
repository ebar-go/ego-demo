// route 使用单独的路由模块来管理中间件与路由
package route

import (
	"ego-demo/http/handler"
	"github.com/ebar-go/ego/http/middleware"
	"github.com/gin-gonic/gin"
)

// Load
func Load(router *gin.Engine)  {
	// 加载中间件,trace,cors,requestLog,recover
	router.Use(middleware.Trace, middleware.CORS, middleware.RequestLog, middleware.Recover)

	router.GET("/", handler.IndexHandler)

	// 定义版本，方便版本升级
	v1 := router.Group("v1")
	{
		// 登录
		v1.POST("user/login", handler.UserAuthHandler)

		// 注册
		v1.POST("user/register", handler.UserRegisterHandler)
	}

	// TODO 权限校验
}
