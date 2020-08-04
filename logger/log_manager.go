package logger

import (
	"log"
	"os"
)

type logManager struct {
}

var manager *logManager

func init() {
	manager = &logManager{}
	file, err := os.OpenFile("log.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	log.SetOutput(file)
}

func GetInstance() *logManager {
	return manager
}

func (l *logManager) Err() *logManager {
	log.SetPrefix("[error]")
	return manager
}

func (l *logManager) Info() *logManager {
	log.SetPrefix("[info]")
	return manager
}
