package logging

import "fmt"

func Init() {
	//SetFormatter(logrus.Stand)
	//logrus.Debug()
	fmt.Println("hello i`m from common")
}

//func SetFormatter(logger *logrus.Logger) {
//	logger.SetFormatter(&logrus.JSONFormatter{
//		FieldMap: logrus.FieldMap{},
//	})
//}
