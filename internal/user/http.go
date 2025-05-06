package main

// go-backend/main.go

import (
	"context"
	"github.com/Ellan98/ding-water-service/user/app/command/query"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/Ellan98/ding-water-service/user/app"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	app app.Application
}

func (h HTTPServer) PostChatCompletion(c *gin.Context, model string) {
	key := "PostChatCompletionRequest"
	val, _ := c.Get(key)
	logrus.Debugf("should bind params %+v", val)
	ctx := context.WithValue(c.Request.Context(), "PostChatCompletionRequest", val)
	//*gin.Context 当作普通的 context.Context 使用了，编译没问题，但失去了 *gin.Context 的所有方法，包括 ShouldBind()。
	o, err := h.app.Queries.PostChatCompletion.Handle(ctx, query.PostChatCompletion{Model: model, Key: key})
	if err != nil {
		logrus.Debug("最后一层输出", err)
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	//TODO Something
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    o.Reply,
	})
}
