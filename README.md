# ego-demo
ego demo project

## 说明
本demo主要介绍`github.com/ebar-go/ego`怎么快速且规范的开发go的web服务。   
涉及内容：配置文件，gin的web服务，mysql服务，请求接收，数据校验器，响应输出，中间件等。细看demo与注释

### 数据库
通过db.sql，初始化数据库

### 配置项
根据实际情况，更改自己的配置项。如端口，数据库连接等

### 启动
```
go run main.go
```

### 接口
- 用户注册   
```
// 请求地址
localhost:8081/v1/user/register
// 请求方式
post(json)
// 请求参数
{
	"email":"hongker@github.com",
	"pass":"123456"
}
// 响应参数
{
    "status_code": 0,
    "message": "success",
    "data": null,
    "meta": {
        "trace": {
            "trace_id": "TraceIdeea596a4-ef3c-46ea-90db-2c20e56fc725",
            "request_id": "request:2d892e13-1ab1-4b16-8d55-4284cf38ff4f"
        },
        "pagination": null
    }
}
```

- 用户登录
```
localhost:8081/v1/user/login
// 请求方式
post(json)
// 请求参数
{
	"email":"hongker@github.com",
	"pass":"123456"
}
```
