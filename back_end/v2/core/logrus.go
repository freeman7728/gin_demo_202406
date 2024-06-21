/*
 * @Description:
 * @author: freeman7728
 * @Date: 2024-04-05 14:11:45
 * @LastEditTime: 2024-04-05 15:01:02
 * @LastEditors: freeman7728
 */
package core

import (
	"bytes"
	"database_lesson/global"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var leveoColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		leveoColor = gray
	case logrus.WarnLevel:
		leveoColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		leveoColor = red
	default:
		leveoColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	log := global.Config.Logger
	timestamp := entry.Time.Format("2006-01-02 -> 15:03:04")
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s : %d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, leveoColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, leveoColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout)
	mLog.SetReportCaller(global.Config.Logger.Show_line)
	mLog.SetFormatter(&LogFormatter{})
	mLog.SetLevel(logrus.DebugLevel)
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)
	InitDefaultLogger()
	return mLog
}

func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Config.Logger.Show_line)
	logrus.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
