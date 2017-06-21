package utils

import (
	"github.com/Sirupsen/logrus"
	"os"
	"path"
	"time"
	"fmt"
	"syscall"
)

//for log
type Logger struct {
	*logrus.Logger
}

var(
	logPath = path.Join(os.Getenv("HOME"),"raft","log")
)

func NewLogger(name string) *Logger {
	fname := path.Join(logPath,fmt.Sprintf("%s-%s",name,time.Now().Format("2006-01-02")))
	f, err := os.OpenFile(fname, syscall.O_WRONLY | syscall.O_APPEND | syscall.O_CREAT, 0666)
	if(err != nil){
		panic(err)
		return nil
	}

	log := &logrus.Logger{
		Out:       f,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}

	return &Logger{
		log,
	}
}


