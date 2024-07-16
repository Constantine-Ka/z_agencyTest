package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {

	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		ForceColors:   true,
		FullTimestamp: true,
		PadLevelText:  true,
	}

	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}
	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	errFile, err := os.OpenFile("logs/err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	warnFile, err := os.OpenFile("logs/warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	infoFile, err := os.OpenFile("logs/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}
	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})
	l.AddHook(&writerHook{
		Writer:    []io.Writer{errFile},
		LogLevels: []logrus.Level{logrus.ErrorLevel},
	})
	l.AddHook(&writerHook{
		Writer:    []io.Writer{warnFile},
		LogLevels: []logrus.Level{logrus.WarnLevel},
	})
	l.AddHook(&writerHook{
		Writer:    []io.Writer{infoFile},
		LogLevels: []logrus.Level{logrus.InfoLevel},
	})
	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
