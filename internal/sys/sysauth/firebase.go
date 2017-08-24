package sysauth

import (
	"context"
	"errors"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2/google"
	"gopkg.in/zabawaba99/firego.v1"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/sys/sysconst"
)

// InitFirebase initialize the Firebase client and set the Firebase instance in the dao package
func InitFirebase() (err error) {
	// Load the file containing Firebase secret
	secretFile, ok := os.LookupEnv(sysconst.FirebaseSecret)
	if !ok || secretFile == "" {
		err = errors.New("Env " + sysconst.FirebaseSecret + " not found")
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

	dao.SetFirebaseClient(firego.New("https://ionline-17da9.firebaseio.com/", conf.Client(context.Background())))

	return
}
