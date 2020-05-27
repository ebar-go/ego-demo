package route

import (
	"ego-demo/http/handler"
	"github.com/ebar-go/ego/http/middleware"
	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine)  {
	// 加载中间件
	router.Use(middleware.Trace, middleware.CORS, middleware.RequestLog)

	router.GET("/", handler.IndexHandler)
}
