package xlog

var logInst *logger

func Init(level string) {
	logInst = newLogger(level)
}

func Debug(format string, v ...interface{}) {
	logInst.doLog(DEBUG, format, v...)
}

func Info(format string, v ...interface{}) {
	logInst.doLog(INFO, format, v...)
}

func Error(format string, v ...interface{}) {
	logInst.doLog(ERROR, format, v...)
}

func Fatal(format string, v ...interface{}) {
	logInst.doLog(FATAL, format, v...)
}
