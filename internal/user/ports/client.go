package ports

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GinServerOptions struct {
	BaseURL      string
	Middlewares  any
	ErrorHandler any
}

type ServerInterface interface {

	// (POST /customer/{customerID}/orders)
	// PostCustomerCustomerIDOrders(c *gin.Context, customerID string)

	//chat/{model}/completion
	PostChatCompletion(c *gin.Context, model string)
	// Get(c *gin.Context, problem string)completion
}

type ServerInterfaceWrapper struct {
	Handler ServerInterface
	// HandlerMiddlewares []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// func (siw *ServerInterfaceWrapper) GetDeepSeekAnswer(c *gin.Context) {
// 	// 从路径中获取参数
// 	problem := c.Param("problem")

// 	// 可选：你可以加一行校验是否为空
// 	if problem == "" {
// 		siw.ErrorHandler(c, errors.New("Missing or empty path parameter: problem"), http.StatusBadRequest)
// 		return
// 	}

// 	// 调用业务逻辑
// 	siw.Handler.GetDeepSeekAnswer(c, problem)
// }

func (siw *ServerInterfaceWrapper) PostChatCompletion(c *gin.Context) {

	//Param 查询 路径参数
	model := c.Param("model")
	//Query 查询？后参数
	// id := c.Query("id")
	//PostForm 查询 post请求携带数据
	// c.PostForm("id")
	var postChatCompletionRequest PostChatCompletionRequest
	if err := c.ShouldBind(&postChatCompletionRequest); err != nil {
		siw.ErrorHandler(c, errors.New("Missing or empty  parameter: postChatCompletionRequest"), http.StatusBadRequest)
		return
	}
	logrus.Debugf("query should bind params %+v", postChatCompletionRequest)
	//在上下文中临时存储自定义键值对，以便在请求的生命周期中 使用。 后续通过Get获取
	c.Set("PostChatCompletionRequest", postChatCompletionRequest)

	if model == "" {
		siw.ErrorHandler(c, errors.New("Missing or empty path parameter: model"), http.StatusBadRequest)
		return
	}

	siw.Handler.PostChatCompletion(c, model)
}

func RegisterHandlersWithOptions(router *gin.Engine, server ServerInterface, options GinServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: server,
		// HandlerMiddlewares: options.Middlewares,
		// ErrorHandler:       errorHandler,
	}
	// fmt.Print(wrapper)
	// router.POST(options.BaseURL+"/customer/:customerID/orders", wrapper.PostCustomerCustomerIDOrders)
	router.POST(options.BaseURL+"/chat/:model/completion", wrapper.PostChatCompletion)
}

// {chat_session_id: "5c949655-39cb-4219-98c7-2ac39df533b9", parent_message_id: null,…}
// chat_session_id
// :
// "5c949655-39cb-4219-98c7-2ac39df533b9"
// parent_message_id
// :
// null
// prompt
// :
// "种自己的花，爱自己的宇宙"
// ref_file_ids
// :
// []
// search_enabled
// :
// false
// thinking_enabled
// :
// false
