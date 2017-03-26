package xlog

import (
	"os"
	"os/signal"
	"syscall"
)

func handleSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR2)
	for {
		<-c

		if logInst.level == DEBUG {
			instance.doLog(INFO, "log level changed to INFO")
			logInst.level = INFO
		} else {
			instance.doLog(INFO, "log level changed to DEBUG")
			logInst.level = DEBUG
		}
	}
}
