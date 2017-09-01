package testrest

import (
	"github.com/maprost/assertion"
	"github.com/maprost/restclient"
	"github.com/maprost/restclient/rctest"

	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/test/testutil"
)

// TestRoute assert that Get : /test-route returns the expected result
func TestRoute(assert assertion.Assert, auth types.AuthChecksum, expected restclient.Result) restclient.ResponseItem {
	result := restclient.Get(test.ServerURL()+testutil.TestRouteClientURL).AddHeader("Cookie", auth.Cookie()).NoLogger().SendAndGetResponseItem()
	rctest.AssertResult(assert, result.Result, expected)
	return result
}

// TestRoute204 assert that Get : /test-route returns a 200
func TestRoute204(assert assertion.Assert, auth types.AuthChecksum) restclient.ResponseItem {
	return TestRoute(assert, auth, rctest.Status204())
}
