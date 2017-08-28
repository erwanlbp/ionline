package dao

import (
	"encoding/json"
	"fmt"

	"errors"

	"github.com/erwanlbp/ionline/internal/data/dao/internal"
	"github.com/erwanlbp/ionline/internal/data/host"
	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/util"
)

// Serie represents a Serie in Firebase
type Serie struct {
	ID          string         `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Host        host.Host      `json:"host,omitempty"`
	Quality     types.Quality  `json:"quality,omitempty"`
	Language    types.Language `json:"language,omitempty"`
	Season      int            `json:"season,omitempty"`
	LastEpisode int            `json:"episode,omitempty"`
	URL         string         `json:"url,omitempty"`
}

// SerieStringFormat is the format to describe a serie
const SerieStringFormat = "{id:%v name:%v host:%v quality:%v language:%v season:%v lastEpisode:%v url:%v}"

// Constants describing Firebase paths
const (
	pathSeries = "series"
)

// FillFromJSON parse a serie from a json byte slice
func (serie *Serie) FillFromJSON(serieBytes []byte) (err error) {
	err = json.Unmarshal(serieBytes, serie)

	// Check required fields
	_, okLang := types.LanguageName[serie.Language]
	if !okLang {
		err = errors.New("Field language isn't valid")
	}

	_, okQual := types.QualityName[serie.Quality]
	if !okQual {
		err = errors.New("Field quality isn't valid")
	}

	return
}

// FindAllSeries returns all the series at /series in Firebase
func FindAllSeries(log logging.Logger) (series []Serie, err error) {
	// Get datas from Firebase
	seriesFB := make(map[string]Serie)
	err = internal.Firebase.Child(pathSeries).Value(&seriesFB)
	internal.LogValue(log, pathSeries)
	if err != nil {
		return
	}

	// Create the slice of Serie
	for id, serie := range seriesFB {
		serie.ID = id
		series = append(series, serie)
	}

	return
}

// Push a serie in Firebase
func (serie *Serie) Push(log logging.Logger) (err error) {
	pushed, err := internal.Firebase.Child(pathSeries).Push(serie)
	internal.LogPush(log, serie, pathSeries)
	if err != nil {
		return
	}

	serie.ID = util.ParseID(pushed.URL())

	return
}

// Delete a serie in Firebase
func (serie *Serie) Delete(log logging.Logger) (err error) {
	err = internal.Firebase.Child(pathSeries).Child(serie.ID).Remove()
	internal.LogRemove(log, pathSeries, serie.ID)
	return
}

// String describe the object
func (serie *Serie) String() string {
	return fmt.Sprintf(SerieStringFormat,
		serie.ID,
		serie.Name,
		host.HostName[serie.Host],
		types.QualityName[serie.Quality],
		types.LanguageName[serie.Language],
		serie.Season,
		serie.LastEpisode,
		serie.URL)
}
