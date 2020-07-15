package main

import (
	"ego-demo/http/route"
	"fmt"
	"github.com/ebar-go/ego"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/component/event"
	"github.com/ebar-go/egu"
)

func init() {
	// 加载配置
	egu.SecurePanic(app.Config().LoadFile("app.yaml"))

	// 初始化数据库
	egu.SecurePanic(app.InitDB())

	egu.SecurePanic(app.Redis().Connect())

	// 支持停止http服务时的回调
	event.Listen(event.BeforeHttpShutdown, func(ev event.Event) {
		// 关闭数据库
		fmt.Println("close database")
		_ = app.DB().Close()
	})
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name hongker
// @contact.url http://hongker.github.io
// @contact.email xiaok2013@live.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func main() {
	s := ego.HttpServer()

	// 加载路由
	route.Load(s.Router)

	// 启动
	egu.SecurePanic(s.Start())
}
