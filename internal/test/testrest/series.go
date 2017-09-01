package testrest

import (
	"github.com/maprost/assertion"
	"github.com/maprost/restclient"
	"github.com/maprost/restclient/rctest"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/test"
)

// GetSeries assert that Get : /series returns the expected result
func GetSeries(assert assertion.Assert, auth types.AuthChecksum, expected restclient.Result) (output string) {
	output, result := restclient.Get(test.ServerURL()+urlpath.SeriesClientURL()).AddHeader("Cookie", auth.Cookie()).NoLogger().SendAndGetResponse()
	rctest.AssertResult(assert, result, expected)
	return output
}

// GetSeries200 assert that Get : /series returns a 200
func GetSeries200(assert assertion.Assert, auth types.AuthChecksum) string {
	return GetSeries(assert, auth, rctest.Status200())
}

// AddSerie assert that Post : /series returns the expected result
func AddSerie(assert assertion.Assert, serie dao.Serie, auth types.AuthChecksum, expected restclient.Result) {
	result := restclient.Post(test.ServerURL()+urlpath.AddSerieClientURL()).AddJsonBody(&serie).AddHeader("Cookie", auth.Cookie()).NoLogger().Send()
	rctest.AssertResult(assert, result, expected)
}

// AddSerie204 assert that Post : /series returns a 204
func AddSerie204(assert assertion.Assert, serie dao.Serie, auth types.AuthChecksum) {
	AddSerie(assert, serie, auth, rctest.Status204())
}

// DeleteSerie assert that Delete : /series returns the expected result
func DeleteSerie(assert assertion.Assert, id string, auth types.AuthChecksum, expected restclient.Result) {
	result := restclient.Delete(test.ServerURL()+urlpath.DeleteSerieClientURL(id)).AddHeader("Cookie", auth.Cookie()).NoLogger().Send()
	rctest.AssertResult(assert, result, expected)
}

// DeleteSerie204 assert that Delete : /series returns a 204
func DeleteSerie204(assert assertion.Assert, id string, auth types.AuthChecksum) {
	DeleteSerie(assert, id, auth, rctest.Status204())
}
