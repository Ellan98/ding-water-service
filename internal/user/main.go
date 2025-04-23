package main

import (
	"fmt"

	"github.com/Ellan98/ding-water-service/common/config"
	"github.com/Ellan98/ding-water-service/common/logging"
	"github.com/Ellan98/ding-water-service/common/server"
	"github.com/Ellan98/ding-water-service/user/ports"
	"github.com/Ellan98/ding-water-service/user/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	logging.Init()
	if err := config.NewViperConfig(); err != nil {
		fmt.Errorf("loading config file error : %s", err)
	}
}

func main() {
	serviceName := viper.GetString("user.service-name")
	fmt.Printf("current service name : %s \n", serviceName)
	// 1.service/application.go
	// 2.app/app.go  command/query.go
	// 3.common/decorator/query.go   logging.go
	// 4.app/command/query.go 的 handle
	// 5.http.go
	application := service.NewApplication()

	//application := app.NewApplication()

	//deepSeek api
	// router.GET("/ping", chatHandler)

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{
			app: application,
		}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})

}
