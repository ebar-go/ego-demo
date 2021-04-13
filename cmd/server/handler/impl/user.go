package impl

import (
	"ego-demo/internal/dto/request"
	"ego-demo/internal/enum"
	"ego-demo/internal/service"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/egu"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

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
func (handler userHandler) Auth(ctx *gin.Context) {
	// 通过结构体获取参数
	var req request.UserAuthRequest

	// 校验参数
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 使用抛出异常的方式，截断代码逻辑，让recover输出响应内容，减少return
		egu.SecurePanic(errors.New(enum.InvalidParam, err.Error()))
	}

	// 调用service的Auth方法，获得结果
	res, err := handler.userService.Auth(req)

	// 有错就抛panic
	egu.SecurePanic(err)

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
func (handler userHandler) Register(ctx *gin.Context) {
	var req request.UserRegisterRequest

	// 校验参数
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 使用抛出异常的方式，截断代码逻辑，让recover输出响应内容，减少return
		egu.SecurePanic(errors.New(enum.InvalidParam, err.Error()))
	}

	// 调用service的Auth方法，获得结果
	err := handler.userService.Register(req)

	// 有错就抛panic
	egu.SecurePanic(err)

	// 输出响应内容
	response.WrapContext(ctx).Success(nil)

}
