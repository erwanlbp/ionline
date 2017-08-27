package testrest

import (
	"github.com/maprost/assertion"
	"github.com/maprost/restclient"
	"github.com/maprost/restclient/rctest"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/test"
)

// GetSeries assert that Get : /series returns the expected result
func GetSeries(assert assertion.Assert, expected restclient.Result) (output string) {
	output, result := restclient.Get(test.ServerURL() + urlpath.SeriesClientURL()).NoLogger().SendAndGetResponse()
	rctest.AssertResult(assert, result, expected)
	return output
}

// GetSeries200 assert that Get : /series returns a 200
func GetSeries200(assert assertion.Assert) string {
	return GetSeries(assert, rctest.Status200())
}

// AddSerie assert that Post : /series returns the expected result
func AddSerie(assert assertion.Assert, serie dao.Serie, expected restclient.Result) {
	result := restclient.Post(test.ServerURL() + urlpath.AddSerieClientURL()).AddJsonBody(&serie).NoLogger().Send()
	rctest.AssertResult(assert, result, expected)
}

// AddSerie204 assert that Post : /series returns a 204
func AddSerie204(assert assertion.Assert, serie dao.Serie) {
	AddSerie(assert, serie, rctest.Status204())
}

// DeleteSerie assert that Delete : /series returns the expected result
func DeleteSerie(assert assertion.Assert, id string, expected restclient.Result) {
	result := restclient.Delete(test.ServerURL() + urlpath.DeleteSerieClientURL(id)).NoLogger().Send()
	rctest.AssertResult(assert, result, expected)
}

// DeleteSerie204 assert that Delete : /series returns a 204
func DeleteSerie204(assert assertion.Assert, id string) {
	DeleteSerie(assert, id, rctest.Status204())
}
