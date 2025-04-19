package main

import (
	"fmt"
	"github.com/Ellan98/ding-water-service/common/config"
	"github.com/Ellan98/ding-water-service/common/logging"
)

func init() {
	logging.Init()
	config.NewViperConfig()
}

func main() {
	fmt.Println("hello world")
}
