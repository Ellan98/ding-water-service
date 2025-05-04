package query

import (
	"context"

	"github.com/Ellan98/ding-water-service/common/decorator"
	"github.com/Ellan98/ding-water-service/user/domain"
	"github.com/sirupsen/logrus"
)

// 请求模型
type PostChatCompletion struct {
	Model string
	Key   string
}

type PostChatCompletionHandler decorator.QueryHandler[PostChatCompletion, *domain.User]

// 应用服务层 依赖于 domain.Repository 接口
type postChatCompletionHandler struct {
	userRepo domain.Repository
}

// 在service 文件中 进行注入
func NewPostChatCompletionHandler(
	userRepo domain.Repository,
	logger *logrus.Entry,
) PostChatCompletionHandler {
	if userRepo == nil {
		panic("nil userRepo")
	}
	return decorator.ApplyQueryDecorators[PostChatCompletion, *domain.User](postChatCompletionHandler{userRepo: userRepo}, logger)
}

// 先调用日志 handle 再调用 这里的handle,  在调用 user_inmem_repository 的get
func (g postChatCompletionHandler) Handle(ctx context.Context, query PostChatCompletion) (*domain.User, error) {
	u, err := g.userRepo.Post(ctx, query.Model, query.Key)
	logrus.Debug("------", u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
