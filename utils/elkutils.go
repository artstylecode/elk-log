package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
)

// Debug 输出debug日志
func Debug(logData interface{}, appName string, module string) {
	log(logData, appName, module, logrus.DebugLevel)
}

// Info 输出info日志
func Info(logData interface{}, appName string, module string) {
	log(logData, appName, module, logrus.InfoLevel)
}

// Error 输出error日志
func Error(logData interface{}, appName string, module string) {
	log(logData, appName, module, logrus.ErrorLevel)
}

// Warn 输出Warn日志
func Warn(logData interface{}, appName string, module string) {
	log(logData, appName, module, logrus.WarnLevel)
}

// Fatal 输出Fatal日志
func Fatal(logData interface{}, appName string, module string) {
	log(logData, appName, module, logrus.FatalLevel)
}

// Panic 输出Panic日志
func Panic(logData interface{}, appName string, module string) {
	log(logData, appName, module, logrus.PanicLevel)
}

// 公有日志输出基础类
func log(logData interface{}, appName string, module string, level logrus.Level) {
	formatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "class",
		},
	}
	if logData == nil {
		logData = make(map[string]interface{})
	}
	message := ""
	if reflect.TypeOf(logData).Name() == "string" {
		message = logData.(string)
		logData = make(map[string]interface{})
	}
	// logrus设置
	logrus.SetFormatter(formatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	entry := logrus.WithFields(logrus.Fields{
		"appName":     appName,
		"module":      module,
		"requestId":   "123456",
		"remoteIp":    "",
		"traceId":     "",
		"spanId":      "",
		"parent":      "",
		"thread":      "",
		"stack_trace": "",
		"line":        -1,
		"data":        logData,
	})

	if level == logrus.DebugLevel {
		entry.Debug(message)
	} else if level == logrus.InfoLevel {
		entry.Info(message)

	} else if level == logrus.WarnLevel {
		entry.Warn(message)
	} else if level == logrus.ErrorLevel {
		entry.Error(message)
	} else if level == logrus.FatalLevel {
		entry.Fatal(message)
	} else if level == logrus.PanicLevel {
		entry.Panic(message)
	}
}
