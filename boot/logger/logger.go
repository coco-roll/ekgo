package logger

import (
	"ekgo/api/boot/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

var Log = logrus.New()

//日志初始化
func Load(logPath string) *logrus.Logger {

	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		log.Println("日志启动异常" + err.Error())
	}

	Log.Out = src

	Log.SetLevel(logrus.InfoLevel)

	//文件最大保存时间
	maximum := config.Get.Log.Maximum
	//文件最大保存时间
	split := config.Get.Log.Split

	logWriter, err := rotatelogs.New(
		logPath+"%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(logPath),                            // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(maximum)*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(split)*time.Hour), // 日志切割时间间隔
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})

	Log.AddHook(lfHook)

	return Log
}
