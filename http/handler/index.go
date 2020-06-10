package handler

import (
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
)

// IndexHandler
func IndexHandler(ctx *gin.Context) {
	response.WrapContext(ctx).Success(nil)
}
