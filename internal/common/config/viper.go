package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

func init() {
	if err := NewViperConfig(); err != nil {
		panic(err)
	}
}

var once sync.Once

/**
命名返回值(err error ) 变量err的作用域覆盖整个函数体
*/

func NewViperConfig() (err error) {
	//once.Do 保证内部的函数只执行一次，且执行过程是线程安全的。
	once.Do(func() {
		err = newViperConfig()
	})
	return
}

func newViperConfig() error {
	relPath, err := getRelativePathFromCaller()
	if err != nil {
		return err
	}
	viper.SetConfigName("global")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(relPath)
	viper.EnvKeyReplacer(strings.NewReplacer("_", "-"))

	return viper.ReadInConfig()

}

func getRelativePathFromCaller() (relPath string, err error) {
	callerPwd, err := os.Getwd()
	fmt.Printf("current path %s\n", callerPwd)
	if err != nil {
		return
	}
	_, here, _, _ := runtime.Caller(0)
	relPath, err = filepath.Rel(callerPwd, filepath.Dir(here))
	fmt.Printf("caller from %s, here: %s ,relpath: %s\n", callerPwd, here, relPath)
	return

}
