package test

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/maprost/assertion"
	"github.com/maprost/restclient"
	"github.com/maprost/restclient/rctest"
	"gopkg.in/zabawaba99/firego.v1"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/extdep"
	"github.com/erwanlbp/ionline/internal/handler"
	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/sys/sysauth"
	"github.com/erwanlbp/ionline/internal/sys/sysconst"
	"github.com/erwanlbp/ionline/internal/sys/systime"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/test/testmock"
	"github.com/erwanlbp/ionline/internal/test/testutil"
)

var server *httptest.Server

var firebaseInitialized bool
var commonExtdepMocked bool
var googleAuthExtdepMocked bool

// InitInternal is the smaller step of the tests initialization
// Use it for tests that doesn't have external dependecy
func InitInternal(t *testing.T) assertion.Assert {
	// If the common extdep has been mocked in a previous test, we unmock it in this one
	if commonExtdepMocked {
		extdep.CommonClient = &extdep.CommonImpl{}
		commonExtdepMocked = false
	}

	systime.Reset()

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
func InitRest(t *testing.T) (assertion.Assert, string, types.AuthChecksum) {
	assert, unikey := Init(t)

	// Initialize the server
	if server == nil {
		router := handler.Init()
		testutil.InitTestRoutes(router)
		server = httptest.NewServer(router)
	}

	if !googleAuthExtdepMocked {
		MockGoogleAuthExtdep(&testmock.MockedGoogleAuth{})
	}

	authChecksum := LoginAndCreateAuthChecksum(assert)

	return assert, unikey, authChecksum
}

// InitRestAndLogger is the same as InitRest() + returns a Logger
func InitRestAndLogger(t *testing.T) (assertion.Assert, string, types.AuthChecksum, logging.Logger) {
	assert, unikey, authChecksum := InitRest(t)
	log := logging.NewLogger()
	return assert, unikey, authChecksum, log
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

// LoginAndCreateAuthChecksum create a valid AuthChecksum and return it
func LoginAndCreateAuthChecksum(assert assertion.Assert) (authChecksum types.AuthChecksum) {

	// Call the login page for the state
	// Can't use testrest function cause of cycles
	responseItem := restclient.Get(ServerURL() + urlpath.LoginPageClientURL()).NoLogger().SendAndGetResponseItem()
	rctest.AssertResult(assert, responseItem.Result, rctest.Status200())

	cookies, ok := responseItem.Header("Set-Cookie")
	assert.True(ok, "Can't get cookie "+sysconst.AuthCookieName+" from ", urlpath.LoginPageClientURL(), " : ", responseItem.Error())

	for _, cookie := range cookies {
		authChecksum, ok = types.ExtractAuthChecksum(cookie)
		if ok {
			break
		}
	}
	assert.NotEqual(authChecksum, "", "Can't retrieve cookie "+sysconst.AuthCookieName+" from ", urlpath.LoginPageClientURL(), " request")

	// Call the auth page for login the test user
	// Can't use testrest function cause of cycles
	responseItem = restclient.Get(ServerURL()+urlpath.AuthClientURL()).
		AddQueryParam(urlpath.StateQueryParam, authChecksum).
		AddQueryParam(urlpath.CodeQueryParam, testmock.CodeAuthTest).
		AddHeader("Cookie", authChecksum.Cookie()).
		NoLogger().
		SendAndGetResponseItem()
	rctest.AssertResult(assert, responseItem.Result, rctest.Status200())

	assert.Nil(responseItem.Error(), "Error during ", urlpath.AuthClientURL(), " request : ", responseItem.Error())

	return
}

// ===== MOCK =====

// MockCommonExtdep replace the extdep.CommonClient by the one received
func MockCommonExtdep(mock extdep.Common) {
	extdep.CommonClient = mock
	commonExtdepMocked = true
}

// MockGoogleAuthExtdep replace the extdep.GoogleAuthClient by the one received
func MockGoogleAuthExtdep(mock extdep.GoogleAuth) {
	extdep.GoogleAuthClient = mock
	googleAuthExtdepMocked = true
}
