package dao_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/data/host"
	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/test/testinject"
)

// ----- Serie FillFromJSON -----

func TestSerie_FillFromJSON(t *testing.T) {
	assert := test.InitInternal(t)

	serie := dao.Serie{
		ID:          "testId",
		Name:        "testName",
		Host:        host.ZoneTelechargementWs,
		Quality:     types.Quality720,
		Language:    types.LangVOSTFR,
		Season:      1,
		LastEpisode: 2,
		URL:         "http://serie.test",
	}
	serieJSON, err := json.Marshal(serie)
	assert.Nil(err)

	var serieUnmarshal dao.Serie
	err = serieUnmarshal.FillFromJSON(serieJSON)
	assert.Nil(err)

	assert.Equal(serieUnmarshal, serie)
}

func TestSerie_FillFromJSON_BadQuality(t *testing.T) {
	assert := test.InitInternal(t)

	serie := dao.Serie{
		ID:          "testId",
		Name:        "testName",
		Host:        host.ZoneTelechargementWs,
		Quality:     types.Quality(-1),
		Language:    types.LangVOSTFR,
		Season:      1,
		LastEpisode: 2,
		URL:         "http://serie.test",
	}
	serieJSON, err := json.Marshal(serie)
	assert.Nil(err)

	var serieUnmarshal dao.Serie
	err = serieUnmarshal.FillFromJSON(serieJSON)
	assert.NotNil(err)
}

func TestSerie_FillFromJSON_BadLanguage(t *testing.T) {
	assert := test.InitInternal(t)

	serie := dao.Serie{
		ID:          "testId",
		Name:        "testName",
		Host:        host.ZoneTelechargementWs,
		Quality:     types.Quality720,
		Language:    types.Language(-1),
		Season:      1,
		LastEpisode: 2,
		URL:         "http://serie.test",
	}
	serieJSON, err := json.Marshal(serie)
	assert.Nil(err)

	var serieUnmarshal dao.Serie
	err = serieUnmarshal.FillFromJSON(serieJSON)
	assert.NotNil(err)
}

// ----- Serie Push -----

func TestSerie_Push(t *testing.T) {
	assert, unikey, log := test.InitAndLogger(t)

	serie := dao.Serie{
		Name:        unikey,
		Quality:     types.Quality720,
		Language:    types.LangVOSTFR,
		Host:        host.ZoneTelechargementWs,
		Season:      1,
		LastEpisode: 2,
		URL:         "http://" + unikey + ".test",
	}
	err := serie.Push(log)
	assert.Nil(err)

	// Test if it has been inserted
	series, err := dao.FindAllSeries(log)
	assert.Nil(err)
	assert.Contains(series, serie)
}

func TestSerie_Push_FirebaseError(t *testing.T) {
	assert, unikey, log := test.InitAndLogger(t)

	test.WantFirebaseError()

	serie := dao.Serie{
		Name:        unikey,
		Quality:     types.Quality720,
		Language:    types.LangVOSTFR,
		Host:        host.ZoneTelechargementWs,
		Season:      1,
		LastEpisode: 2,
		URL:         "http://" + unikey + ".test",
	}
	err := serie.Push(log)
	assert.NotNil(err)
}

// ----- Serie Delete -----

func TestSerie_Delete(t *testing.T) {
	assert, unikey, log := test.InitAndLogger(t)

	serie := testinject.PushSerie(assert, unikey, log)

	err := serie.Delete(log)
	assert.Nil(err)

	// Test if it has been deleted
	series, err := dao.FindAllSeries(log)
	assert.Nil(err)
	for _, serieFetched := range series {
		assert.NotEqual(serieFetched.ID, serie.ID)
	}
}

func TestSerie_Delete_FirebaseError(t *testing.T) {
	assert, unikey, log := test.InitAndLogger(t)

	serie := testinject.PushSerie(assert, unikey, log)

	test.WantFirebaseError()

	err := serie.Delete(log)
	assert.NotNil(err)
}

// ----- Serie FindAll -----

func TestFindAllSeries(t *testing.T) {
	assert, unikey, log := test.InitAndLogger(t)

	serie1 := testinject.PushSerie(assert, unikey, log)
	serie2 := testinject.PushSerie(assert, unikey, log)

	series, err := dao.FindAllSeries(log)
	assert.Nil(err)
	assert.Contains(series, serie1)
	assert.Contains(series, serie2)
}

func TestFindAllSeries_FirebaseError(t *testing.T) {
	assert, unikey, log := test.InitAndLogger(t)

	testinject.PushSerie(assert, unikey, log)

	test.WantFirebaseError()

	_, err := dao.FindAllSeries(log)
	assert.NotNil(err)
}

// ----- Serie String -----

func TestSerie_String(t *testing.T) {
	assert := test.InitInternal(t)

	serie := dao.Serie{
		ID:          "testId",
		Name:        "testName",
		Quality:     types.Quality720,
		Language:    types.LangVOSTFR,
		Host:        host.ZoneTelechargementWs,
		Season:      1,
		LastEpisode: 2,
		URL:         "http://string.test",
	}
	expected := fmt.Sprintf(dao.SerieStringFormat,
		serie.ID,
		serie.Name,
		host.HostName[serie.Host],
		types.QualityName[serie.Quality],
		types.LanguageName[serie.Language],
		serie.Season,
		serie.LastEpisode,
		serie.URL,
	)

	assert.Equal(serie.String(), expected)
}
