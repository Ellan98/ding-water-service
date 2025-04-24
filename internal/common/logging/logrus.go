package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func Init() {
	SetFormatter(logrus.StandardLogger())
	fmt.Println("hello i`m from common")
	logrus.SetLevel(logrus.DebugLevel)
}

func SetFormatter(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyMsg:   "message",
		},
	})
	//strconv.ParseBool： 将字符串 转换为 bool值
	//如果isLocal 为true 则属于开发环境 os.Getenv  获取环境变量
	if isLocal, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV")); isLocal {

	}
}
