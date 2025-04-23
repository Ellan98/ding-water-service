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
		Problem: "hello world ",
	})
	return &MemoryUserRepository{
		lock:  &sync.RWMutex{},
		store: s,
	}
}

func (m MemoryUserRepository) Get(ctx context.Context, problem string) (*domain.User, error) {
	for i, v := range m.store {
		logrus.Infof("m.store[%d] = %+v", i, v)
	}
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, o := range m.store {
		if o.Problem != "" {
			return o, nil
		}
	}
	return nil, domain.NotFound{Problem: problem}

}
