package main

import (
	"fmt"
	"github.com/Ellan98/ding-water-server/common/config"
)

func init() {
	if err := config.NewViperConfig(); err != nil {

	}
}

func main() {
	fmt.Println("hello world")
}
