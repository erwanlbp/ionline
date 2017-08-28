package webcrawler_test

import (
	"testing"

	"fmt"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/data/host"
	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/test/testmock"
	"github.com/erwanlbp/ionline/internal/webcrawler"
)

func TestActualizeSerie(t *testing.T) {
	assert, log := test.InitInternalAndLogger(t)

	mockedExtdep := &testmock.MockedCommonExtdep{}
	test.MockCommonExtdep(mockedExtdep)

	serie := &dao.Serie{
		Quality:  types.Quality720,
		Language: types.LangVOSTFR,
		URL:      "http://" + host.ZoneTelechargementWsString + "/exclus/26986-telecharger-game-of-thrones-saison-7-french-hdtv-streaming.html",
	}

	err := webcrawler.ActualizeSerie(log, serie)
	assert.Nil(err)
}

func TestActualizeSerie_WrongHost(t *testing.T) {
	assert, log := test.InitInternalAndLogger(t)

	serie := &dao.Serie{
		Quality:  types.Quality720,
		Language: types.LangVOSTFR,
		URL:      "http://wrong.host.test/something",
	}

	err := webcrawler.ActualizeSerie(log, serie)
	assert.NotNil(err)
}

func TestActualizeSerie_SeasonNotFound(t *testing.T) {
	assert, log := test.InitInternalAndLogger(t)

	mockedExtdep := &testmock.MockedCommonExtdep{}
	test.MockCommonExtdep(mockedExtdep)

	serie := &dao.Serie{
		Quality:  types.Quality720,
		Language: types.LangVOSTFR,
		URL:      "http://" + host.ZoneTelechargementWsString + "/zt-wrong-season",
	}

	err := webcrawler.ActualizeSerie(log, serie)
	fmt.Println(serie)
	assert.NotNil(err)
}

func TestActualizeSerie_EpisodeNotFound(t *testing.T) {
	assert, log := test.InitInternalAndLogger(t)

	mockedExtdep := &testmock.MockedCommonExtdep{}
	test.MockCommonExtdep(mockedExtdep)

	serie := &dao.Serie{
		Quality:  types.Quality720,
		Language: types.LangVOSTFR,
		URL:      "http://" + host.ZoneTelechargementWsString + "/zt-wrong-episode",
	}

	err := webcrawler.ActualizeSerie(log, serie)
	assert.NotNil(err)
}
