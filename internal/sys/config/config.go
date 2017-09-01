package config

import (
	"flag"
	"os"
	"path/filepath"
)

var serverHost string
var logger string
var projectPath string
var firebaseCredentials string
var googleAuth string

func init() {
	flag.StringVar(&serverHost, "host", "http://localhost:8080", "The host of the server")
	flag.StringVar(&logger, "log", "stdoutcache", "The type of logger")
	flag.StringVar(&projectPath, "projectpath", os.Getenv("GOPATH")+"/src/github.com/erwanlbp/ionline/", "Path to the project directory")
	flag.StringVar(&firebaseCredentials, "firebase-credentials", "IONLINE_SECRET_FIREBASE", "Name of the environment variable for the Firebase Credentials")
	flag.StringVar(&googleAuth, "google-auth", "IONLINE_SECRET_GOOGLE_AUTH", "Name of the environment variable for the Google Authentication")
	flag.Parse()

	// If the public directory is not accessible, no need to start the server
	path, err := filepath.Abs(projectPath)
	if err == nil {
		_, err = os.Stat(path)
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

// ProjectPath returns the path to the public folder
func ProjectPath() string {
	return projectPath
}

// FirebaseCredentials returns the name of the environment variable for the Firebase Credentials
func FirebaseCredentials() string {
	return firebaseCredentials
}

// GoogleAuth returns the name of the environment variable for the Firebase Authentication
func GoogleAuth() string {
	return googleAuth
}
