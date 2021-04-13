package main

import (
	"ego-demo/cmd/server/handler"
	"ego-demo/cmd/server/route"
	"ego-demo/internal/repository"
	"ego-demo/internal/service"
	"github.com/ebar-go/ego"
	"log"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name hongker
// @contact.url http://hongker.github.io
// @contact.email xiaok2013@live.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	app := ego.App()
	// 加载配置
	if err := app.LoadConfig("configs/app.yaml"); err != nil {
		log.Fatalf("unable to load config: %v", err)
	}

	// 注入repository
	repository.Inject(app.Container())
	// 注入service
	service.Inject(app.Container())
	// 注入handler
	handler.Inject(app.Container())

	// 加载路由
	if err := app.LoadRouter(route.Loader); err != nil {
		log.Fatalf("unable to load router: %v", err)
	}

	// 启动http服务
	app.ServeHTTP()
	app.Run()
}
