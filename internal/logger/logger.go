package logger

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

func InitLogger(logLevel string) *logrus.Logger {
	log := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Ошибка при парсинге уровня логирования: %v", err)
	}
	log.SetLevel(level)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return fmt.Sprintf("%s:%d", frame.Function, frame.Line), ""
		},
	})

	return log
}
