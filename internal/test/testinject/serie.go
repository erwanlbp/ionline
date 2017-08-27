package testinject

import (
	"github.com/maprost/assertion"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/data/host"
	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/sys/logging"
)

// PushSerie insert a Serie and test everything went well
func PushSerie(assert assertion.Assert, unikey string, log logging.Logger) dao.Serie {
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
	assert.NotEqual(serie.ID, "")

	return serie
}
