package helper

import (
	"ego-demo/pkg/service/data"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/http/middleware"
	"github.com/ebar-go/ego/utils/conv"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/gin-gonic/gin"
)

// GeLoginUserFromContext 通过Context获取登录的用户信息
func GeLoginUserFromContext(ctx *gin.Context) data.User {
	claims := middleware.GetCurrentClaims(ctx)
	if claims == nil {
		secure.Panic(errors.Unauthorized("please login first"))
	}

	userClaims := map[string]interface{}(claims)

	var user data.User
	if err := conv.Map2Struct(userClaims["user"].(map[string]interface{}), &user); err != nil {
		secure.Panic(errors.Unauthorized("please login first"))
	}

	return user
}
