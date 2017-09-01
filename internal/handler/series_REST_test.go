package handler_test

import (
	"strings"
	"testing"

	"github.com/maprost/restclient/rctest"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/data/host"
	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/handler"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/test/testinject"
	"github.com/erwanlbp/ionline/internal/test/testrest"
)

func TestSeriesListTemplateData_DeleteSerieURL(t *testing.T) {
	assert := test.InitInternal(t)

	tmplt := handler.SeriesListTemplateData{}

	url := tmplt.DeleteSerieURL("testId")
	assert.Equal(url, urlpath.DeleteSerieClientURL("testId"))
}

func TestListSeries(t *testing.T) {
	assert, unikey, authChecksum, log := test.InitRestAndLogger(t)

	serie := testinject.PushSerie(assert, unikey, log)

	output := testrest.GetSeries200(assert, authChecksum)
	assert.True(strings.Contains(output, serie.ID))
}

func TestListSeries_FirebaseError(t *testing.T) {
	assert, _, authChecksum := test.InitRest(t)

	test.WantFirebaseError()

	testrest.GetSeries(assert, authChecksum, rctest.Status500())
}

func TestAddSerie(t *testing.T) {
	assert, unikey, authChecksum := test.InitRest(t)

	serie := dao.Serie{
		Name:     unikey,
		Host:     host.ZoneTelechargementWs,
		Quality:  types.QualityHDTV,
		Language: types.LangVF,
		URL:      "https://www.zone-telechargement.ws/exclus/26986-telecharger-game-of-thrones-saison-7-french-hdtv-streaming.html",
	}

	testrest.AddSerie204(assert, serie, authChecksum)
}

func TestAddSerie_MissRequiredFields(t *testing.T) {
	assert, unikey, authChecksum := test.InitRest(t)

	serie := dao.Serie{
		Name:     unikey,
		Host:     host.ZoneTelechargementWs,
		Quality:  types.Quality(-1),
		Language: types.Language(-1),
		URL:      "http://" + unikey + ".test",
	}

	testrest.AddSerie(assert, serie, authChecksum, rctest.Status400())
}

func TestAddSerie_WrongURL(t *testing.T) {
	assert, unikey, authChecksum := test.InitRest(t)

	serie := dao.Serie{
		Name:     unikey,
		Host:     host.ZoneTelechargementWs,
		Quality:  types.Quality720,
		Language: types.LangVOSTFR,
		URL:      "http://" + unikey + ".test",
	}

	testrest.AddSerie(assert, serie, authChecksum, rctest.Status400())
}

func TestAddSerie_FirebaseError(t *testing.T) {
	assert, unikey, authChecksum := test.InitRest(t)

	serie := dao.Serie{
		Name:     unikey,
		Host:     host.ZoneTelechargementWs,
		Quality:  types.QualityHDTV,
		Language: types.LangVF,
		URL:      "https://www.zone-telechargement.ws/exclus/26986-telecharger-game-of-thrones-saison-7-french-hdtv-streaming.html",
	}

	test.WantFirebaseError()

	testrest.AddSerie(assert, serie, authChecksum, rctest.Status500())
}

func TestDeleteSerie(t *testing.T) {
	assert, unikey, authChecksum, log := test.InitRestAndLogger(t)

	serie := testinject.PushSerie(assert, unikey, log)

	testrest.DeleteSerie204(assert, serie.ID, authChecksum)

	series, err := dao.FindAllSeries(log)
	assert.Nil(err)

	for _, serieFetched := range series {
		assert.NotEqual(serieFetched.ID, serie.ID)
	}
}

func TestDeleteSerie_FirebaseError(t *testing.T) {
	assert, unikey, authChecksum, log := test.InitRestAndLogger(t)

	serie := testinject.PushSerie(assert, unikey, log)

	test.WantFirebaseError()

	testrest.DeleteSerie(assert, serie.ID, authChecksum, rctest.Status500())
}
