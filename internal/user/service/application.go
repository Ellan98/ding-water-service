package service

import (
	"github.com/Ellan98/ding-water-service/user/adapters"
	"github.com/Ellan98/ding-water-service/user/app"
	"github.com/Ellan98/ding-water-service/user/app/command/query"
	"github.com/sirupsen/logrus"
)

func NewApplication() app.Application {
	return newApplication()
}

func newApplication() app.Application {

	userRepo := adapters.NewMemoryUserRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	return app.Application{
		Queries: app.Queries{
			PostChatCompletion: query.NewPostChatCompletionHandler(userRepo, logger),
		},
	}
}
