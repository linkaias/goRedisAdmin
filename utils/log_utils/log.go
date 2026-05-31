package log_utils

import (
	"fmt"
	"goRedisAdmin/global/initData"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// logClient is the shared logrus instance used across the application.
var logClient *logrus.Logger

// logDataStruct represents one asynchronous log event.
type logDataStruct struct {
	error  error
	title  string
	result interface{}
}

// waitWriteLog is the buffered queue for asynchronous log writes.
var waitWriteLog = make(chan *logDataStruct, 1000)

// init configures logrus output, rotation strategy and JSON formatter hook.
func init() {
	cfg := initData.IniRead.Section("log")

	logClient = logrus.New()

	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println("err", err)
		return
	}
	// Disable default stdout output; logs are persisted via hooks.
	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)

	logPath := cfg.Key("log_path").String()
	logDay, _ := cfg.Key("max_save_day").Int()

	logWriter, err := rotatelogs.New(
		logPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(logPath),                          // Symlink to current log file.
		rotatelogs.WithMaxAge(24*time.Hour*time.Duration(logDay)), // Maximum retention duration.
		rotatelogs.WithRotationTime(time.Hour*24),                 // Daily rotation interval.
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
	}
	// Persist structured logs as JSON for easier downstream parsing.
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})

	logClient.AddHook(lfHook)
}

// RunLog starts background workers to consume log queue and write log entries.
func RunLog() {
	for i := 0; i < 2; i++ {
		go func() {
			for log := range waitWriteLog {
				// Write info log when no error attached; otherwise write error log.
				if log.error == nil {
					logClient.Infof("[ title: %s ; body: %+v ]",
						log.title,
						log.result,
					)
				} else {
					logClient.Errorf("[ title: %s ;error: %s; body: %+v ]",
						log.title,
						log.error.Error(),
						log.result,
					)
				}
			}
		}()
	}
}

// WriteLog enqueues one log record for asynchronous persistence.
func WriteLog(title string, err error, v interface{}) {
	waitWriteLog <- &logDataStruct{
		error:  err,
		title:  title,
		result: v,
	}
}
