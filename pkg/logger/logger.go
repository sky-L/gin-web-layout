package logger

import (
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func InitLog() {
	rotate := &lumberjack.Logger{
		Filename:   "log.Filename",
		MaxSize:    500, // 最大的文件2000M
		MaxBackups: 30,  // 最多保留10个文件
		MaxAge:     10,  // 最长保留7天
	}
	logrus.SetOutput(rotate)
}
