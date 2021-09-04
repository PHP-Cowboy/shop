package log_source

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init()  {
	logConf := LoadLogConfig()
	file,err := os.OpenFile(logConf.LogDir,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0755)
	if err != nil {
		panic(err)
	}
	//defer file.Close()

	Logger.Out = file
	levelMap := map[string]logrus.Level{
		"info":logrus.InfoLevel,
	}
	Logger.SetLevel(levelMap[logConf.LogLevel])
	Logger.SetFormatter(&logrus.TextFormatter{})
}