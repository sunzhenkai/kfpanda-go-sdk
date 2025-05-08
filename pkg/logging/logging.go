package logging

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.Formatter = &logrus.TextFormatter{
		DisableQuote: true,
	}
	Logger.SetLevel(logrus.InfoLevel)
	ll := &lumberjack.Logger{
		Filename:   "log/kfpanda-client.log",
		MaxSize:    256,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   true,
	}
	Logger.Out = ll
}
