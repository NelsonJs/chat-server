package logger

import (
	"io"
	"log"
	"os"
)

type logManager struct {
	log_ *log.Logger
}

var manager *logManager

func init() {
	manager = &logManager{}
	file, err := os.OpenFile("log.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	log_s := log.New(io.MultiWriter(file, os.Stderr), "--->", log.Ldate|log.Ltime|log.Lshortfile)
	manager.log_ = log_s
}

func GetInstance() *logManager {
	return manager
}

func (l *logManager) ErrLog() *log.Logger {
	l.log_.SetPrefix("[error]")
	return l.log_
}

func (l *logManager) InfoLog() *log.Logger {
	l.log_.SetPrefix("[info]")
	return l.log_
}
