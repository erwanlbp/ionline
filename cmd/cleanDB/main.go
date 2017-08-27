package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2/google"
	"gopkg.in/zabawaba99/firego.v1"
)

var firebase *firego.Firebase

func main() {
	// Init Firebase test database
	fmt.Println("Initializing client ...")
	err := initFirebase()
	if err != nil {
		fmt.Println("ERROR Initializing Firebase:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Cleaning database ...")
	err = firebase.Remove()
	if err != nil {
		fmt.Println("ERROR Removing DB root")
		os.Exit(1)
	}

	fmt.Println("Cleaned Firebase database : ionline-test")
}

func initFirebase() error {
	// Load the file containing Firebase secret
	secretFile, ok := os.LookupEnv("IONLINE_TEST_SECRET_FIREBASE")
	if !ok || secretFile == "" {
		return errors.New("Env IONLINE_TEST_SECRET_FIREBASE not found")
	}

	d, err := ioutil.ReadFile(secretFile)
	if err != nil {
		return err
	}

	conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		return err
	}

	firebase = firego.New("https://ionline-test.firebaseio.com/", conf.Client(context.Background()))

	return nil
}
