package ports

import (
	"github.com/gin-gonic/gin"
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
	GetDeepSeekAnswer(c *gin.Context)
	// Get(c *gin.Context, problem string)
}

type ServerInterfaceWrapper struct {
	Handler ServerInterface
	// HandlerMiddlewares []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

func RegisterHandlersWithOptions(router *gin.Engine, server ServerInterface, options GinServerOptions) {
	// wrapper := ServerInterfaceWrapper{
	// 	Handler: server,
	// 	// HandlerMiddlewares: options.Middlewares,
	// 	// ErrorHandler:       errorHandler,
	// }
	// fmt.Print(wrapper)
	// router.POST(options.BaseURL+"/customer/:customerID/orders", wrapper.PostCustomerCustomerIDOrders)
	router.GET(options.BaseURL+"/ping", server.GetDeepSeekAnswer)
}
