package logger

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func InitLogger(logLevel string) {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Ошибка при парсинге уровня логирования: %v", err)
	}
	log.SetLevel(level)
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return fmt.Sprintf("%s:%d", frame.Function, frame.Line), ""
		},
	})
}
