package bootstrap

import (
	"log"
	"os"
	"sync"
)

var (
	logger     *log.Logger
	loggerOnce sync.Once
)

func InitLogger() {
	loggerOnce.Do(func() {
		logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	})
}

func GetLogger() *log.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}
