package testrest

import (
	"github.com/maprost/assertion"
	"github.com/maprost/restclient"
	"github.com/maprost/restclient/rctest"

	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/test"
)

// GetIndex assert that Get : / returns the expected result
func GetIndex(assert assertion.Assert, auth types.AuthChecksum, expected restclient.Result) restclient.ResponseItem {
	result := restclient.Get(test.ServerURL()+urlpath.IndexClientURL()).AddHeader("Cookie", auth.Cookie()).NoLogger().SendAndGetResponseItem()
	rctest.AssertResult(assert, result.Result, expected)
	return result
}

// GetIndex200 assert that Get : / returns a 200
func GetIndex200(assert assertion.Assert, auth types.AuthChecksum) restclient.ResponseItem {
	return GetIndex(assert, auth, rctest.Status200())
}
