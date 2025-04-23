package ports

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinServerOptions struct {
	BaseURL      string
	Middlewares  any
	ErrorHandler any
}

type ServerInterface interface {

	// (POST /customer/{customerID}/orders)
	// PostCustomerCustomerIDOrders(c *gin.Context, customerID string)

	// (GET /customer/{customerID}/orders/{orderID})
	GetDeepSeekAnswer(c *gin.Context, problem string)
	// Get(c *gin.Context, problem string)
}

type ServerInterfaceWrapper struct {
	Handler ServerInterface
	// HandlerMiddlewares []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

func (siw *ServerInterfaceWrapper) GetDeepSeekAnswer(c *gin.Context) {
	// 从路径中获取参数
	problem := c.Param("problem")

	// 可选：你可以加一行校验是否为空
	if problem == "" {
		siw.ErrorHandler(c, errors.New("Missing or empty path parameter: problem"), http.StatusBadRequest)
		return
	}

	// 调用业务逻辑
	siw.Handler.GetDeepSeekAnswer(c, problem)
}

func RegisterHandlersWithOptions(router *gin.Engine, server ServerInterface, options GinServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: server,
		// HandlerMiddlewares: options.Middlewares,
		// ErrorHandler:       errorHandler,
	}
	// fmt.Print(wrapper)
	// router.POST(options.BaseURL+"/customer/:customerID/orders", wrapper.PostCustomerCustomerIDOrders)
	router.GET(options.BaseURL+"/ping/:problem", wrapper.GetDeepSeekAnswer)
}
