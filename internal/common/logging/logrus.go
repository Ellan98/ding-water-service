package logging

func Init() {
	//SetFormatter(logrus.Stand)
	logrus.Debug()
}

//func SetFormatter(logger *logrus.Logger) {
//	logger.SetFormatter(&logrus.JSONFormatter{
//		FieldMap: logrus.FieldMap{},
//	})
//}
