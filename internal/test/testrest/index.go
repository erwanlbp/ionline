package testrest

import (
	"github.com/maprost/assertion"
	"github.com/maprost/restclient"
	"github.com/maprost/restclient/rctest"

	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/test"
)

// GetIndex assert that Get : / returns the expected result
func GetIndex(assert assertion.Assert, expected restclient.Result) {
	result := restclient.Get(test.ServerURL() + urlpath.IndexClientURL()).NoLogger().Send()
	rctest.AssertResult(assert, result, expected)
}

// GetIndex200 assert that Get : / returns a 200
func GetIndex200(assert assertion.Assert) {
	GetIndex(assert, rctest.Status200())
}
