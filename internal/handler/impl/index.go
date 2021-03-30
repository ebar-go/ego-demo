package impl

import (
	"ego-demo/internal/handler"
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
)

type indexHandler struct {

}

func newIndexHandler() handler.IndexHandler {
	return &indexHandler{}
}


// IndexHandler
func (handler indexHandler) Index(ctx *gin.Context) {
	response.WrapContext(ctx).Success(nil)
}
