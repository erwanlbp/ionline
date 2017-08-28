package config

import (
	"flag"
	"os"
	"path/filepath"
)

var serverHost string
var logger string
var pathPublic string
var firebaseAuth string

func init() {
	flag.StringVar(&serverHost, "host", "http://localhost:8080", "The host of the server")
	flag.StringVar(&logger, "log", "stdoutcache", "The type of logger")
	flag.StringVar(&pathPublic, "public", os.Getenv("GOPATH")+"/src/github.com/erwanlbp/ionline/internal/public/", "Path to the public directory")
	flag.StringVar(&firebaseAuth, "firebase-auth", "IONLINE_SECRET_FIREBASE", "Name of the environment variable for the Firebase Authentication")
	flag.Parse()

	// If the public directory is not accessible, no need to start the server
	path, err := filepath.Abs(pathPublic)
	if err == nil {
		_, err = os.Stat(path)
		os.IsNotExist(err)
	}
	if err != nil {
		panic(err.Error())
	}
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

// PathToPublic returns the path to the public folder
func PathToPublic() string {
	return pathPublic
}

// FirebaseAuth returns the name of the environment variable for the Firebase Authentication
func FirebaseAuth() string {
	return firebaseAuth
}
