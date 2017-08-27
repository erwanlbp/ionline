package test

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/maprost/assertion"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/extdep"
	"github.com/erwanlbp/ionline/internal/handler"
	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/sys/sysauth"
	"gopkg.in/zabawaba99/firego.v1"
)

var server *httptest.Server

var firebaseInitialized bool
var commonExtdepMocked bool

// InitInternal is the smaller step of the tests initialization
// Use it for tests that doesn't have external dependecy
func InitInternal(t *testing.T) assertion.Assert {
	// If the common extdep has been mocked in a previous test, we unmock it in this one
	if commonExtdepMocked {
		extdep.CommonClient = &extdep.CommonImpl{}
		commonExtdepMocked = false
	}

	return assertion.New(t)
}

// InitInternalAndLogger is the same as InitInternal() + returns a Logger
func InitInternalAndLogger(t *testing.T) (assertion.Assert, logging.Logger) {
	assert := InitInternal(t)
	log := logging.NewLogger()
	return assert, log
}

// Init initialize the test environment
// Initialize Firebase connection
// + InitInternal()
func Init(t *testing.T) (assertion.Assert, string) {
	assert := InitInternal(t)

	// Initialize Firebase
	if !firebaseInitialized {
		err := sysauth.InitFirebase()
		assert.Nil(err)
		firebaseInitialized = true
	}

	return assert, uniqueKey(t)
}

// InitAndLogger is the same as Init() + returns a Logger
func InitAndLogger(t *testing.T) (assertion.Assert, string, logging.Logger) {
	assert, unikey := Init(t)
	log := logging.NewLogger()
	return assert, unikey, log
}

// InitRest initialize a test server to test the REST calls
// + Init()
func InitRest(t *testing.T) (assertion.Assert, string) {
	assert, unikey := Init(t)

	// Initialize the server
	if server == nil {
		router := handler.Init()
		server = httptest.NewServer(router)
	}

	return assert, unikey
}

// InitRestAndLogger is the same as InitRest() + returns a Logger
func InitRestAndLogger(t *testing.T) (assertion.Assert, string, logging.Logger) {
	assert, unikey := InitRest(t)
	log := logging.NewLogger()
	return assert, unikey, log
}

// ServerURL returns the URL of the test server
func ServerURL() string {
	if server == nil {
		return "server-not-working"
	}
	return server.URL
}

func uniqueKey(t *testing.T) string {
	return fmt.Sprintf("%v-%v", t.Name(), time.Now().Nanosecond())
}

// WantFirebaseError change the Firebase client to a fake one that will provoke error when used
func WantFirebaseError() {
	fb := firego.New("http://firebase.wrong.test/", nil)
	dao.SetFirebaseClient(fb)
	firebaseInitialized = false
	return
}

// MockCommonExtdep replace the extdep.CommonClient by the one received
func MockCommonExtdep(mock extdep.Common) {
	extdep.CommonClient = mock
	commonExtdepMocked = true
}
