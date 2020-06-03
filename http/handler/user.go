package handler

import (
	"ego-demo/http/helper"
	"ego-demo/pkg/enum/statusCode"
	"ego-demo/pkg/request"
	"ego-demo/pkg/service"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/gin-gonic/gin"
)

// UserAuthHandler 用户登录
func UserAuthHandler(ctx *gin.Context)  {
	// 通过结构体获取参数
	var req request.UserAuthRequest

	// 校验参数
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 使用抛出异常的方式，截断代码逻辑，让recover输出响应内容，减少return
		secure.Panic(errors.New(statusCode.InvalidParam, err.Error()))
	}

	// 调用service的Auth方法，获得结果
	res, err := service.User().Auth(req)

	// 有错就抛panic
	secure.Panic(err)

	// 输出响应内容
	response.WrapContext(ctx).Success(res)

}


// UserRegisterHandler 用户注册
func UserRegisterHandler(ctx *gin.Context)  {
	var req request.UserRegisterRequest

	// 校验参数
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 使用抛出异常的方式，截断代码逻辑，让recover输出响应内容，减少return
		secure.Panic(errors.New(statusCode.InvalidParam, err.Error()))
	}

	// 调用service的Auth方法，获得结果
	err := service.User().Register(req)

	// 有错就抛panic
	secure.Panic(err)

	// 输出响应内容
	response.WrapContext(ctx).Success(nil)

}

// GetUserInfoHandler 获取用户信息
func GetUserInfoHandler(ctx *gin.Context)  {
	loginUser := helper.GeLoginUserFromContext(ctx)
	response.WrapContext(ctx).Success(response.Data{
		"email": loginUser.Email,
	})
}