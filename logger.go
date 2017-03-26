package xlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type logger struct {
	level     uint
	starttime time.Time
	duration  time.Duration
	file      *os.File
	filename  string
	lock      sync.RWMutex
}

var once sync.Once
var instance *logger

func (l *logger) prefix(level string) string {
	pc, fn, line, _ := runtime.Caller(3)
	names := strings.Split(fn, "/")
	name := names[len(names)-1]
	f := runtime.FuncForPC(pc)
	return fmt.Sprintf("[%s]: %s, <%s, %d>, ", level, name, f.Name(), line)
}

func (l *logger) writeLog(prefix string, format string, v ...interface{}) {
	log.Printf(prefix+format+"\n", v...)
}

func (l *logger) renameFile() {
	l.file.Close()
	os.Rename(l.filename, l.filename+".1")

	return
}

func (l *logger) openFile() {
	var exist = true
	var err error
	if _, err = os.Stat(l.filename); os.IsNotExist(err) {
		exist = false
	}

	var perm int
	if exist == true {
		perm = os.O_RDWR
	} else {
		perm = os.O_RDWR | os.O_CREATE
	}

	l.file, err = os.OpenFile(l.filename, perm, 0666)
	if err != nil {
		log.Printf("open file %s failed with perm %d. %s.", l.filename, err.Error(), perm)
		os.Exit(1)
	}

	f, _ := l.file.Stat()
	l.file.Seek(f.Size(), 0)

	log.SetOutput(l.file)
}

func (l *logger) doLog(level uint, format string, v ...interface{}) {
	if level >= l.level {
		/*		l.lock.RLock()
				if time.Time.Sub(time.Now(), l.starttime) > l.duration {
					l.lock.RUnlock()
					l.lock.Lock()
					if time.Time.Sub(time.Now(), l.starttime) > l.duration {
						l.renameFile()

						var err error
						l.file, err = os.Create(l.filename)
						if err != nil {
							fmt.Println(err.Error())
							os.Exit(-1)
						}

						log.SetOutput(l.file)
						l.starttime = time.Now()
					}
					l.lock.Unlock()
					l.lock.RLock()
				}
				l.writeLog(l.prefix(LEVELS[level]), format, v...)
				l.lock.RUnlock()*/
		strPre := l.prefix(LEVELS[level])
		strLog := fmt.Sprintf(format, v...)
		fmt.Printf("%s%s\n", strPre, strLog)
	}
}

func newLogger(level string) *logger {
	once.Do(func() {
		log.SetFlags(log.LstdFlags)

		instance = &logger{}

		switch level {
		case "DEBUG":
			instance.level = DEBUG
		case "INFO":
			instance.level = INFO
		case "ERROR":
			instance.level = ERROR
		case "FATAL":
			instance.level = FATAL
		default:
			instance.level = INFO
		}
		//instance.openFile()

		go handleSignal()
	})
	return instance
}
