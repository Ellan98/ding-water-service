package main

import (
	"fmt"

	"github.com/Ellan98/ding-water-service/common/config"
	"github.com/Ellan98/ding-water-service/common/logging"
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
	fmt.Println("hello world", serviceName)
}
