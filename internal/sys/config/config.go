package config

import "flag"

var serverHost string
var logger string

func init() {
	flag.StringVar(&serverHost, "host", "http://localhost:8080", "The host of the server")
	flag.StringVar(&logger, "log", "stdoutcache", "The type of logger")
}

// ServerHost returns the host of the server
func ServerHost() string {
	return serverHost
}

// ServerPort returns the port of the server
func ServerPort() string {
	return "8080"
}

// Logger returns the type of logger
func Logger() string {
	return logger
}
