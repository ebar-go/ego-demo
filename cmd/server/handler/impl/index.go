package impl

import (
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
)

type indexHandler struct {

}

func NewIndexHandler() *indexHandler {
	return &indexHandler{}
}


// IndexHandler
func (handler indexHandler) Index(ctx *gin.Context) {
	response.WrapContext(ctx).Success(nil)
}
