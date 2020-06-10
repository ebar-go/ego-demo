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
// @Summary 用户登录
// @Description 通过邮箱和密码登录，换取token
// @Accept  json
// @Produce json
// @Param email body string true "邮箱"
// @Param pass body string true "密码"
// @Success 0 "success"
// @Failure 500 "error"
// @Router /user/auth [post]
func UserAuthHandler(ctx *gin.Context) {
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
// @Summary 用户注册
// @Description 通过邮箱和密码注册账户
// @Accept  json
// @Produce json
// @Param req body request.UserRegisterRequest true "请求参数"
// @Success 0 "success"
// @Failure 500 "error"
// @Router /user/register [post]
func UserRegisterHandler(ctx *gin.Context) {
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
func GetUserInfoHandler(ctx *gin.Context) {
	loginUser := helper.GeLoginUserFromContext(ctx)
	response.WrapContext(ctx).Success(response.Data{
		"email": loginUser.Email,
	})
}
