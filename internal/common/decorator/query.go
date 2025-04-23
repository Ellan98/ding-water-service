package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Q, R 此时都是 any 类型 [Q, R any]为简洁写法
type QueryHandler[Q, R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}

// 通过链路调用
func ApplyQueryDecorators[H, R any](handle QueryHandler[H, R], logger *logrus.Entry) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		logger: logger,
		base:   handle,
	}

}
