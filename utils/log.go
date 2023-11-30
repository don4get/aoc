package aoc

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type MyFormatter struct{}

var levelList = []string{
	"panic",
	"fatal",
	"error",
	"warn",
	"info",
	"debug",
	"trace",
}

var logFile = "app.log"

func (mf *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// define supported log levels
	level := levelList[int(entry.Level)]
	strList := strings.Split(entry.Caller.File, "/")
	// get the go file name
	fileName := strList[len(strList)-1]
	b.WriteString(fmt.Sprintf("%28s, %5s, [%s:%d] - %s\n",
		// Custom Time Format
		entry.Time.Format("2006-01-02T15:04:05.9999999Z"),
		level, fileName, entry.Caller.Line, entry.Message))
	return b.Bytes(), nil
}

func ConfigureLogging() *logrus.Logger {
	// read the logfile
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		// create the file
		os.Create(logFile)
	}

	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()

	// Log the messages on console as well as log file
	logger.SetOutput(io.MultiWriter(f, os.Stdout))

	logger.SetLevel(logrus.DebugLevel)
	// Required to get line number
	logger.SetReportCaller(true)

	// Set custom formatter
	logger.SetFormatter(&MyFormatter{})
	return logger
}
