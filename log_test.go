package xlog

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const LOG_FILE = "/tmp/log_test.txt"

func run(c string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", c)
	return cmd.Output()
}

func checkExist(substr string, t *testing.T) {
	c := fmt.Sprintf("cat %s", LOG_FILE)
	out, _ := run(c)
	if !strings.Contains(string(out), substr) {
		t.Fail()
	}
}

func checkNotExist(substr string, t *testing.T) {
	c := fmt.Sprintf("cat %s", LOG_FILE)
	out, _ := run(c)
	if strings.Contains(string(out), substr) {
		t.Fail()
	}
}

func TestInit(t *testing.T) {
	run(fmt.Sprintf("rm -f %s", LOG_FILE))
	Init(LOG_FILE, "DEBUG", 2)

}

func TestDebug(t *testing.T) {
	Debug("debug test")
	checkExist("[DEBUG]", t)
}

func TestInfo(t *testing.T) {
	Info("info test")
	checkExist("[INFO]", t)
}

func TestError(t *testing.T) {
	Error("error test")
	checkExist("[ERROR]", t)
}

func TestFatal(t *testing.T) {
	Fatal("fatal test")
	checkExist("[FATAL]", t)
}

func TestCloseDebug(t *testing.T) {
	run(fmt.Sprintf("kill -USR2 %d", os.Getpid()))
	strLog := "debug not exist"
	Debug(strLog)
	checkNotExist(strLog, t)
}

func TestOpenDebug(t *testing.T) {
	run(fmt.Sprintf("kill -USR2 %d", os.Getpid()))
	strLog := "debug exist"
	Debug(strLog)
	checkExist(strLog, t)
}
