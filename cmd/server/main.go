package main

import (
	"ego-demo/cmd/server/route"
	handlerImpl "ego-demo/internal/handler/impl"
	serviceImpl "ego-demo/internal/service/impl"
	"github.com/ebar-go/ego"
	"github.com/ebar-go/egu"

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
	// 加载配置
	app := ego.App()
	egu.SecurePanic(app.LoadConfig("configs/app.yaml"))

	if err := serviceImpl.Inject(app.Container()); err != nil {
		log.Fatalf("inject service failed: %v\n", err)
	}

	if err := handlerImpl.Inject(app.Container()); err != nil {
		log.Fatalf("inject handler failed: %v\n", err)
	}

	egu.SecurePanic(app.Container().Invoke(route.Loader))

	app.ServeHTTP()
	app.Run()
}
