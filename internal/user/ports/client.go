package ports

import "github.com/gin-gonic/gin"

type GinServerOptions struct {
	BaseURL      string
	Middlewares  any
	ErrorHandler any
}

func RegisterHandlersWithOptions(router *gin.Engine, server *HTTPServer) {

}
