package xlog

const (
	DEBUG = iota
	INFO
	ERROR
	FATAL
)

var LEVELS = [4]string{"DEBUG", "INFO", "ERROR", "FATAL"}
