package server

func RunHTTPServer(serviceName string, wrapper func(router *gin.Engine)) {
	//Sub返回新的Viper实例，表示此实例的子树。Sub对键不区分大小写。
	addr := viper.Sub(serviceName).GetString("http-addr")
	if addr == "" {
		//TODO: Waring log
	}
	RunHTTPServerOnAddr(addr)

}

/*
通用 HTTP Server 启动器
*/
func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {

}
