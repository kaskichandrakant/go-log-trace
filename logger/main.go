package logger

import (
	"log"
	"os"
)

type CustomLogger struct {
	loggers map[string]*log.Logger
	file    *os.File
	flag    int
}

func (lgr CustomLogger) attachGenericLoggers() {
	lgr.AddCustom(PrefixMessage, PrefixMessage)
	lgr.AddCustom(PrefixInfo, PrefixInfo)
	lgr.AddCustom(PrefixErr, PrefixErr)
	lgr.AddCustom(PrefixDebug, PrefixDebug)
}

func (lgr CustomLogger) Info(message string) {
	lgr.loggers[PrefixInfo].Println(message)
}

func (lgr CustomLogger) Err(message string) {
	lgr.loggers[PrefixErr].Println(message)
}

func (lgr CustomLogger) Debug(message string) {
	lgr.loggers[PrefixDebug].Println(message)
}

func (lgr CustomLogger) Print(messageType string, message string) {
	lgr.loggers[messageType].Println(message)
}

func (lgr CustomLogger) AddCustom(name string, prefix string) {
	logger := log.New(lgr.file, AddPrefix(":", prefix), lgr.flag)
	logger.SetOutput(os.Stdout)
	lgr.loggers[name] = logger
}

func NewCustomLogger(fileName string, flag int) (*CustomLogger, error) {
	var file *os.File
	var err error
	if fileName != "" {
		if file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err != nil {
			return nil, err
		}
	}
	cstmLgr := &CustomLogger{file: file, flag: flag, loggers: map[string]*log.Logger{}}
	cstmLgr.attachGenericLoggers()
	return cstmLgr, nil
}
