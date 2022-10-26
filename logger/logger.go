package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-log-slack/goslack"
	"runtime"
)

var logger = logrus.New()

func SetLogLevel(level logrus.Level) {
	logger.Level = level
}

func SetLogFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

func SetLogJsonFormatter() {
	logger.Formatter = &logrus.JSONFormatter{}
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		//entry.Data["file"] = fileInfo(2)
		entry.Debug(args...)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileAddressInfo(2)
		entry.Info(args...)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileAddressInfo(2)
		entry.Warn(args...)
	}
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileAddressInfo(2)
		entry.Error(args...)
		slackLogReq := SlackLogRequest{
			Message: fmt.Sprint(args...),
			File:    fileAddressInfo(2),
			Level:   "error",
		}
		_ = ProcessAndSend(slackLogReq, goslack.Alert, "Error")
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		slackLogReq := SlackLogRequest{
			Message: fmt.Sprint(args...),
			File:    fileAddressInfo(2),
			Level:   "fatal",
		}
		_ = ProcessAndSend(slackLogReq, goslack.Alert, "Fatal")
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileAddressInfo(2)
		entry.Fatal(args...)

	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	slackLogReq := SlackLogRequest{
		Message: fmt.Sprint(args...),
		File:    fileAddressInfo(2),
		Level:   "panic",
	}
	_ = ProcessAndSend(slackLogReq, goslack.Alert, "Panic")
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileAddressInfo(2)
	entry.Panic(args...)
}

func fileAddressInfo(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s:%d", file, line)
}