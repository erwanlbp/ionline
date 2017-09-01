package handler_test

import (
	"testing"
	"time"

	"github.com/maprost/restclient/rctest"

	"github.com/erwanlbp/ionline/internal/sys/sysconst"
	"github.com/erwanlbp/ionline/internal/sys/systime"
	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/test/testrest"
)

func TestLogin_expiredDate(t *testing.T) {
	assert, _, _ := test.InitRest(t)

	// +6 days
	now := systime.Now().Add(time.Hour * 24 * 6)
	systime.InitTime(now)

	authChecksum := test.LoginAndCreateAuthChecksum(assert)

	responseItem := testrest.TestRoute204(assert, authChecksum)
	//no redirect to login
	_, ok := responseItem.Header("Set-Cookie")
	assert.False(ok, "Set-Cookie is set when there was a redirect to login.")

	{
		// +24 hours -1 Second
		systime.InitTime(now.Add(sysconst.AuthChecksumExpire).Add(-time.Second))

		responseItem := testrest.TestRoute204(assert, authChecksum)
		// no redirect to index
		_, ok := responseItem.Header("Set-Cookie")
		assert.False(ok, "Set-Cookie is set when there was a redirect to login.")
	}

	{
		// +24 hours +1 Second
		systime.InitTime(now.Add(sysconst.AuthChecksumExpire).Add(time.Second))

		responseItem := testrest.TestRoute(assert, authChecksum, rctest.Status200())
		// redirect to index -> set-Cookie is inside the header
		_, ok := responseItem.Header("Set-Cookie")
		assert.True(ok, "There should be a redirect to index page.")
	}
}
