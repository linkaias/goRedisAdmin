package log_utils

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"goRedisAdmin/global/initData"
	"os"
	"time"
)

var logClient *logrus.Logger

type logDataStruct struct {
	error  error
	title  string
	result interface{}
}

var waitWriteLog = make(chan *logDataStruct, 1000)

func init() {
	cfg := initData.IniRead.Section("log")

	logClient = logrus.New()

	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println("err", err)
		return
	}
	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)

	logPath := cfg.Key("log_path").String()
	logDay, _ := cfg.Key("max_save_day").Int()

	logWriter, err := rotatelogs.New(
		logPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(logPath),                          // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(24*time.Hour*time.Duration(logDay)), // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour*24),                 // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})

	logClient.AddHook(lfHook)
}

func RunLog() {
	for i := 0; i < 2; i++ {
		go func() {
			for log := range waitWriteLog {
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

func WriteLog(title string, err error, v interface{}) {
	waitWriteLog <- &logDataStruct{
		error:  err,
		title:  title,
		result: v,
	}
}
