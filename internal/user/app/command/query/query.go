package query

import (
	"context"

	"github.com/Ellan98/ding-water-service/common/decorator"
	"github.com/Ellan98/ding-water-service/user/domain"
	"github.com/sirupsen/logrus"
)

// 请求模型
type GetDeepSeekAnswer struct {
	Problem string
}

type GetDeepSeekAnswerHandler decorator.QueryHandler[GetDeepSeekAnswer, *domain.User]

// 应用服务层 依赖于 domain.Repository 接口
type getDeepSeekAnswerHandler struct {
	userRepo domain.Repository
}

// 在service 文件中 进行注入
func NewGetDeepSeekAnswerHandler(
	userRepo domain.Repository,
	logger *logrus.Entry,

) GetDeepSeekAnswerHandler {
	if userRepo == nil {
		panic("nil userRepo")
	}
	return decorator.ApplyQueryDecorators[GetDeepSeekAnswer, *domain.User](getDeepSeekAnswerHandler{userRepo: userRepo}, logger)
}

// 先调用日志 handle 再调用 这里的handle
func (g getDeepSeekAnswerHandler) Handle(ctx context.Context, query GetDeepSeekAnswer) (*domain.User, error) {
	u, err := g.userRepo.Get(ctx, query.Problem)
	if err != nil {
		return nil, err
	}
	return u, nil
}
