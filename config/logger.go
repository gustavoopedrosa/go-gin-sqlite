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

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Debug creates non-formatted debug logs.
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

// Info creates non-formatted info logs.
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

// Warning creates non-formatted warning logs.
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

// Error creates non-formatted error logs.
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Debugf creates format enabled debug logs.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

// Infof creates format enabled info logs.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

// Warningf creates format enabled warning logs.
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

// Errorf creates format enabled error logs.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
