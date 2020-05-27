package main

import (
	"ego-demo/http/route"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/config"
	"github.com/ebar-go/ego/http"
	"github.com/ebar-go/ego/http/validator"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/gin-gonic/gin/binding"
)

func init()  {
	// 加载配置
	secure.Panic(config.ReadFromFile("app.yaml"))

	// 初始化数据库
	secure.Panic(app.InitDB())

	// 设置自定义验证器,支持字段命名
	binding.Validator = new(validator.Validator)
}

func main()  {
	s := http.NewServer()

	// 加载路由
	route.Load(s.Router)

	// 启动
	secure.Panic(s.Start())
}
