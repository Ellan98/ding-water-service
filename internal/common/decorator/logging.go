package decorator

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type queryLoggingDecorator[C, R any] struct {
	logger *logrus.Entry
	base   QueryHandler[C, R]
}

func (q queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (rusult R, err error) {
	/*
			logrus.WithFields  logrus 的链式调用创建带有一个结构化的字段（Fields）的 logger 并自动附加一些字段
			example:
		{
		  "level": "Debug",
		  "msg": "Executing query",
		  "query": "GetUserByID",
		  "query_body": "main.UserQuery{Name:\"Alice\", Age:30}"
		}
	*/
	logger := q.logger.WithFields(logrus.Fields{
		"query":      generateActionName(cmd),
		"query_body": fmt.Sprintf("%#v", cmd),
	})
	logger.Debug("Executing query")
	//开发小技巧 流转这个函数后 再进行一次打印
	defer func() {
		if err != nil {
			logger.Info("Executing query successfully")
		} else {
			logger.Error("Failed to executing query", err)
		}
	}()
	return q.base.Handle(ctx, cmd)
}

func generateActionName(cmd any) string {
	return strings.Split(fmt.Sprintf("%T", cmd), ".")[1]
}
