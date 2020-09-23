package handler

import (
	"ego-demo/pkg/service"
	"github.com/ebar-go/ego/http/response"
	"github.com/ebar-go/egu"
	"github.com/gin-gonic/gin"
)

func GetOrderHandler(ctx *gin.Context)  {
	id := egu.String2Int(ctx.Param("id"))
	item, err := service.Order().Get(id)
	egu.SecurePanic(err)
	response.WrapContext(ctx).Success(item)
}
