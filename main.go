package main

import (
	"ego-demo/http/route"
	"github.com/ebar-go/ego/config"
	"github.com/ebar-go/ego/http"
	"github.com/ebar-go/ego/utils/secure"
)

func init()  {
	secure.Panic(config.ReadFromFile("app.yaml"))
}

func main()  {
	s := http.NewServer()

	// 加载路由
	route.Load(s.Router)

	// 启动
	secure.Panic(s.Start())
}
