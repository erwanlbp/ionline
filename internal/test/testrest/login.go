package testrest

import (
	"github.com/maprost/assertion"
	"github.com/maprost/restclient"
	"github.com/maprost/restclient/rctest"

	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/test/testmock"
)

// Login200 returns the header to check for cookies
func Login200(assert assertion.Assert) restclient.ResponseItem {
	return Login(assert, rctest.Status200())
}

// Login returns the header to check for cookies
func Login(assert assertion.Assert, expectedResult restclient.Result) restclient.ResponseItem {
	responseItem := restclient.Get(test.ServerURL() + urlpath.LoginPageClientURL()).NoLogger().SendAndGetResponseItem()
	rctest.AssertResult(assert, responseItem.Result, expectedResult)

	return responseItem
}

// Auth200 tries to authenticate the authChecksum
func Auth200(assert assertion.Assert, authState types.AuthChecksum) restclient.ResponseItem {
	return Auth(assert, authState, testmock.CodeAuthTest, rctest.Status200())
}

// Auth tries to authenticate the authChecksum
func Auth(assert assertion.Assert, authChecksum types.AuthChecksum, code string, expectedResult restclient.Result) restclient.ResponseItem {
	responseItem := restclient.Get(test.ServerURL()+urlpath.AuthClientURL()).
		AddQueryParam(urlpath.StateQueryParam, authChecksum).
		AddQueryParam(urlpath.CodeQueryParam, code).
		AddHeader("Cookie", authChecksum.Cookie()).
		NoLogger().
		SendAndGetResponseItem()
	rctest.AssertResult(assert, responseItem.Result, expectedResult)

	return responseItem
}
