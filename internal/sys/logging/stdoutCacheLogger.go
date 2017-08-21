package logging

import (
	"fmt"
	"os"
	"strings"
)

type stdoutCacheLogger struct {
	logs        []string
	hasWarning  bool
	hasError    bool
	hasCritical bool
}

// NewStdoutCacheLogger returns an instance of a logger that cache logs and print them only when Release is called
func NewStdoutCacheLogger() ExtendedLogger {
	return &stdoutCacheLogger{}
}

// ----- Logger Interface -----

func (l *stdoutCacheLogger) Println(v ...interface{}) {
	s := "[info]  " + fmt.Sprintln(v...)
	l.logs = append(l.logs, s)
}

func (l *stdoutCacheLogger) Printf(format string, v ...interface{}) {
	l.logs = append(l.logs, fmt.Sprintf("[info]  "+format, v))
}

func (l *stdoutCacheLogger) Warnln(v ...interface{}) {
	s := "[warn]  " + fmt.Sprintln(v...)
	l.logs = append(l.logs, s)
	l.hasWarning = true
}

func (l *stdoutCacheLogger) Warnf(format string, v ...interface{}) {
	l.logs = append(l.logs, fmt.Sprintf("[warn]  "+format, v))
	l.hasWarning = true
}

// ----- Extended Logger Interface -----

func (l *stdoutCacheLogger) Error(v ...interface{}) {
	s := "[error] " + fmt.Sprintln(v...)
	l.logs = append(l.logs, s)
	l.hasError = true
}

func (l *stdoutCacheLogger) Critical(v ...interface{}) {
	s := "[fatal] " + fmt.Sprintln(v...)
	l.logs = append(l.logs, s)
	l.hasCritical = true
}

func (l *stdoutCacheLogger) Release() error {
	// Add the clientID, join the logs in one string, send the log
	logStr := strings.Join(l.logs, "")

	if l.hasError {
		fmt.Fprintln(os.Stderr, logStr)
	} else {
		fmt.Fprintln(os.Stdout, logStr)
	}
	return nil
}
