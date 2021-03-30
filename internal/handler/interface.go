package handler

import "github.com/gin-gonic/gin"

type IndexHandler interface {
	Index(ctx *gin.Context)
}

type UserHandler interface {
	Auth(ctx *gin.Context)
	Register(ctx *gin.Context)
}