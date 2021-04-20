package mw

import (
	"ego-demo/internal/entity"
	"errors"
	"github.com/ebar-go/ego/http/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLoginUser(ctx *gin.Context) entity.User {
	claims, exist := middleware.GetClaims(ctx)
	if !exist {
		panic(errors.New("401"))
	}

	userId, _ := strconv.Atoi(claims["id"].(string))
	return entity.User{
		Id:    userId,
		Email: "",
	}
}
