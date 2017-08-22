package logging

import (
	"strings"

	"github.com/erwanlbp/ionline/internal/sys/config"
)

// Logger is the basics functions a Logger should have in the project
type Logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
	Warnln(...interface{})
	Warnf(string, ...interface{})
}

// ExtendedLogger add functions to the Logger interface
// Such as severity and Release()
type ExtendedLogger interface {
	Logger
	Error(...interface{})
	Critical(...interface{})
	Release() error
}

// logType represents a type of logger
type logType int

// Constants for the differents logger type
const (
	stdout = logType(iota)
	stdoutCache
)

// Constants for the strings of the log types
const (
	stdoutString      = "stdout"
	stdoutCacheString = "stdoutCache"
)

// logTypeStringIDMap map the string of a log type to their ID
var logTypeStringIDMap = map[string]logType{
	strings.ToLower(stdoutString):      stdout,
	strings.ToLower(stdoutCacheString): stdoutCache,
}

func getLogType() (logType, bool) {
	lt, ok := logTypeStringIDMap[config.Logger()]
	return lt, ok
}

// NewLogger returns a Logger depending on the logtype received
func NewLogger() ExtendedLogger {
	lt, ok := getLogType()
	if !ok {
		return NewStdoutLogger()
	}
	switch lt {
	case stdoutCache:
		return NewStdoutCacheLogger()
	default:
		return NewStdoutLogger()
	}
}
