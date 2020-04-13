package xheylog

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

type Xlog struct {
	LogLevel int
}

const (
	DEBUG = iota
	INFO
	ERROR
	FATAL
)

var (
	logLevel map[string]int
)

func init() {
	log.SetFlags(log.LstdFlags)
	logLevel = make(map[string]int)
	logLevel["DEBUG"] = DEBUG
	logLevel["INFO"] = INFO
	logLevel["ERROR"] = ERROR
	logLevel["FATAL"] = FATAL
}

func (h *Xlog) dolog(level string, format string, v ...interface{}) {
	if h.LogLevel > logLevel[level] {
		return
	}
	pc, fn, line, _ := runtime.Caller(2)
	names := strings.Split(fn, "/")
	name := names[len(names)-1]
	f := runtime.FuncForPC(pc)
	logTime := time.Now().Format("2006/01/02 15:04:05")
	prefix := fmt.Sprintf("%s [%s]: %s, <%s, %d>, ", logTime, level, name, f.Name(), line)
	strlog := fmt.Sprintf(prefix+format, v...)

	fmt.Println(strlog)
}

func (h *Xlog) Debug(format string, v ...interface{}) {
	h.dolog("DEBUG", format, v...)
}

func (h *Xlog) Info(format string, v ...interface{}) {
	h.dolog("INFO", format, v...)
}

func (h *Xlog) Error(format string, v ...interface{}) {
	h.dolog("ERROR", format, v...)
}

func (h *Xlog) Fatal(format string, v ...interface{}) {
	h.dolog("FATAL", format, v...)
}
