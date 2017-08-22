package logging

import (
	"fmt"
)

type stdoutLogger struct{}

// NewStdoutLogger returns an instance of a StdoutLogger
func NewStdoutLogger() ExtendedLogger {
	return &stdoutLogger{}
}

// ----- Logger Interface -----

func (l *stdoutLogger) Println(v ...interface{}) {
	fmt.Println("[info]\t", fmt.Sprint(v...))
}

func (l *stdoutLogger) Printf(format string, v ...interface{}) {
	fmt.Printf("[info]\t "+format, v)
}

func (l *stdoutLogger) Warnln(v ...interface{}) {
	fmt.Println("[warn]\t", fmt.Sprint(v...))
}

func (l *stdoutLogger) Warnf(format string, v ...interface{}) {
	fmt.Printf("[warn]\t "+format, v)
}

// ----- Extended Logger Interface -----

func (l *stdoutLogger) Error(v ...interface{}) {
	fmt.Println("[error]\t", fmt.Sprint(v...))
}

func (l *stdoutLogger) Critical(v ...interface{}) {
	fmt.Println("[critical]\t", fmt.Sprint(v...))
}

func (l stdoutLogger) Release() error {
	fmt.Println()
	return nil
}
