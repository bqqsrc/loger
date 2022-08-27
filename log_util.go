package loger

import (
	"fmt"
	"log"
	"os"
)

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

var level = LevelTrace
var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

//获取日志等级
func Level() int {
	return level
}

//设置日志等级
func SetLevel(l int) {
	level = l
}

//设置日志输出器
func SetLogger(l *log.Logger) {
	logger = l
}

func Trace(v ...interface{}) {
	if level <= LevelTrace {
		logger.Printf("[T] %v\n", v)
	}
}

func Tracef(format string, v ...interface{}) {
	if level <= LevelTrace {
		msg := fmt.Sprintf(format, v...)
		Trace(msg)
	}
}

func Debug(v ...interface{}) {
	if level <= LevelDebug {
		logger.Printf("[D] %v\n", v)
	}
}

func Debugf(format string, v ...interface{}) {
	if level <= LevelDebug {
		msg := fmt.Sprintf(format, v...)
		Debug(msg)
	}
}

func Info(v ...interface{}) {
	if level <= LevelInfo {
		logger.Printf("[I] %v\n", v)
	}
}

func Infof(format string, v ...interface{}) {
	if level <= LevelInfo {
		msg := fmt.Sprintf(format, v...)
		Info(msg)
	}
}

func Warn(v ...interface{}) {
	if level <= LevelWarning {
		logger.Printf("[W] %v\n", v)
	}
}

func Warnf(format string, v ...interface{}) {
	if level <= LevelWarning {
		msg := fmt.Sprintf(format, v...)
		Warn(msg)
	}
}

func Error(v ...interface{}) {
	if level <= LevelError {
		logger.Printf("[E] %v\n", v)
	}
}

func Errorf(format string, v ...interface{}) {
	if level <= LevelError {
		msg := fmt.Sprintf(format, v...)
		Error(msg)
	}
}

func Critical(v ...interface{}) {
	if level <= LevelCritical {
		logger.Printf("[C] %v\n", v)
	}
}

func Criticalf(format string, v ...interface{}) {
	if level <= LevelCritical {
		msg := fmt.Sprintf(format, v...)
		Critical(msg)
	}
}
