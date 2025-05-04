package adapters

import (
	"context"
	"sync"

	"github.com/Ellan98/ding-water-service/user/domain"
	"github.com/sirupsen/logrus"
)

type MemoryUserRepository struct {
	lock  *sync.RWMutex
	store []*domain.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	s := make([]*domain.User, 0)
	s = append(s, &domain.User{
		Model: "hello world ",
	})
	return &MemoryUserRepository{
		lock:  &sync.RWMutex{},
		store: s,
	}
}

// 考虑 这个方向 构造 deepSeek 请求
func (m MemoryUserRepository) Post(ctx context.Context, model string) (*domain.User, error) {
	for i, v := range m.store {
		logrus.Infof("m.store[%d] = %+v", i, v)
	}
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, o := range m.store {
		if o.Model != "" {
			return o, nil
		}
	}
	return nil, domain.NotFound{Model: model}

}
