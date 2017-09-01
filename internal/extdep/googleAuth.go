package extdep

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/sys/config"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
)

// ----- Interface declaration -----

// GoogleAuth is the extdep inferface for the Google Authentication
type GoogleAuth interface {
	InitGoogleAuthCredentials() error
	GetLoginURL(string) string
	GoogleAuthenticate(string) (dao.User, error)
}

// GoogleAuthClient is the interface to Google authentify
var GoogleAuthClient GoogleAuth = googleAuthImpl{}

type googleAuthImpl struct{}

// ----- Interface implementation -----

type googleAuthCredentials struct {
	Web struct {
		Cid     string `json:"client_id"`
		Csecret string `json:"client_secret"`
	} `json:"web"`
}

var gac *googleAuthCredentials
var confGoogleAuth *oauth2.Config

// InitGoogleAuthCredentials create a config for the Google Auth
func (r googleAuthImpl) InitGoogleAuthCredentials() error {
	// ----- Read and store the application credentials -----
	envVarGoogleAuthCredential, ok := os.LookupEnv(config.GoogleAuth())
	if !ok {
		return errors.New("Environment variable " + config.GoogleAuth() + " is needed for the user authentication")
	}
	file, err := ioutil.ReadFile(envVarGoogleAuthCredential)
	if err != nil {
		return err
	}
	json.Unmarshal(file, &gac)

	// ----- Create a conf variable to call the Google Auth API -----
	confGoogleAuth = &oauth2.Config{
		ClientID:     gac.Web.Cid,
		ClientSecret: gac.Web.Csecret,
		RedirectURL:  config.ServerHost() + urlpath.AuthClientURL(),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	return nil
}

// GetLoginURL return the URL to access the Google Auth page
func (r googleAuthImpl) GetLoginURL(state string) string {
	if confGoogleAuth == nil {
		r.InitGoogleAuthCredentials()
	}
	return confGoogleAuth.AuthCodeURL(state)
}

// GoogleAuthenticate connect to google with the code received and return informations about the connected user
func (r googleAuthImpl) GoogleAuthenticate(code string) (user dao.User, err error) {
	// Validate code
	tok, err := confGoogleAuth.Exchange(context.Background(), code)
	if err != nil {
		return
	}

	// Get user infos
	client := confGoogleAuth.Client(context.Background(), tok)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	userByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(userByte, &user)

	return
}
