package sysauth

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2/google"
	"gopkg.in/zabawaba99/firego.v1"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/sys/config"
)

type firebaseSecret struct {
	ProjectID string `json:"project_id"`
}

// InitFirebase initialize the Firebase client and set the Firebase instance in the dao package
func InitFirebase() (err error) {
	// Load the file containing Firebase secret
	secretFile, ok := os.LookupEnv(config.FirebaseCredentials())
	if !ok || secretFile == "" {
		err = errors.New("Env " + config.FirebaseCredentials() + " not found")
		return
	}

	d, err := ioutil.ReadFile(secretFile)
	if err != nil {
		return
	}

	conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		return
	}

	var fbSec firebaseSecret
	err = json.Unmarshal(d, &fbSec)
	if err != nil {
		return
	}

	dao.SetFirebaseClient(firego.New("https://"+fbSec.ProjectID+".firebaseio.com/", conf.Client(context.Background())))

	return
}
