package testmock

import (
	"errors"

	"github.com/erwanlbp/ionline/internal/data/dao"
)

// Constants for mocking google and authentify a test user
const (
	EmailTest    = "ionline-test@gmail.com"
	CodeAuthTest = "ionline-mocked-code-from-google-auth"
)

// MockedGoogleAuth mocks the Google Authentication
type MockedGoogleAuth struct {
	initialized bool
}

// InitGoogleAuthCredentials fake the initialization
func (r MockedGoogleAuth) InitGoogleAuthCredentials() error {
	r.initialized = true
	return nil
}

// GetLoginURL return a fake URL and initialize if needed
func (r MockedGoogleAuth) GetLoginURL(state string) string {
	if !r.initialized {
		r.InitGoogleAuthCredentials()
	}
	return "http://loginURLWithState:" + state
}

// GoogleAuthenticate authenticate a user if he as the CodeAuthTest
func (r MockedGoogleAuth) GoogleAuthenticate(code string) (user dao.User, err error) {
	if code != CodeAuthTest {
		err = errors.New("Bad CodeAuthTest : " + code)
		return
	}
	user.Email = EmailTest
	return
}
