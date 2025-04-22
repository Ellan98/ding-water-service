package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RunHTTPServer(serviceName string, wrapper func(router *gin.Engine)) {
	//Sub返回新的Viper实例，表示此实例的子树。Sub对键不区分大小写。
	addr := viper.Sub(serviceName).GetString("http-addr")
	if addr == "" {
		//TODO: Waring log
		panic("empth http-addr ")
	}
	RunHTTPServerOnAddr(addr, wrapper)

}

/*
通用 HTTP Server 启动器
*/
func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	//
	apiRouter := gin.New()
	wrapper(apiRouter)
	apiRouter.Group("/api")
	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
	apiRouter.Run(addr)
}
