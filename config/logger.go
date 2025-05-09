package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)
	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (logger *Logger) Debug(v ...interface{}) {
	logger.debug.Println(v...)
}

func (logger *Logger) Info(v ...interface{}) {
	logger.info.Println(v...)
}

func (logger *Logger) Warning(v ...interface{}) {
	logger.warning.Println(v...)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.err.Println(v...)
}

func (logger *Logger) DebugFormated(format string, v ...interface{}) {
	logger.debug.Printf(format, v...)
}

func (logger *Logger) InfoFormated(format string, v ...interface{}) {
	logger.info.Printf(format, v...)
}

func (logger *Logger) WarningFormated(format string, v ...interface{}) {
	logger.warning.Printf(format, v...)
}

func (logger *Logger) ErrorFormated(format string, v ...interface{}) {
	logger.err.Printf(format, v...)
}
