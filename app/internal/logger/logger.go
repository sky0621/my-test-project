package logger

import (
	"fmt"
	"runtime"
)

type AppLogger interface {
	Log(msg string)
}

type appLogger struct {
}

func NewAppLogger() AppLogger {
	return &appLogger{}
}

func (l *appLogger) Log(msg string) {
	fmt.Printf("[Path:%s] %s\n", fileWithLineNum(), msg)
}

func fileWithLineNum() string {
	_, name, line, ok := runtime.Caller(2)
	if !ok {
		return "-"
	}
	return fmt.Sprintf("%s:%d", name, line)
}
