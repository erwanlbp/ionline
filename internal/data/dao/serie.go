package dao

import (
	"encoding/json"

	"github.com/erwanlbp/ionline/internal/data/dao/internal"
	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/util"
)

// Serie represents a Serie in Firebase
type Serie struct {
	ID          string         `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Quality     types.Quality  `json:"quality,omitempty"`
	Language    types.Language `json:"language,omitempty"`
	Season      int            `json:"season,omitempty"`
	LastEpisode int            `json:"episode,omitempty"`
	URL         string         `json:"url,omitempty"`
}

// Constants describing Firebase paths
const (
	pathSeries = "series"
)

// ParseJSON a serie from a byte slice
func (serie *Serie) ParseJSON(serieBytes []byte) (err error) {
	err = json.Unmarshal(serieBytes, serie)
	return
}

// Push a serie in Firebase
func (serie *Serie) Push(log logging.Logger) (err error) {
	pushed, err := internal.Firebase.Child(pathSeries).Push(serie)
	internal.LogPush(log, pushed, serie)
	if err != nil {
		return
	}

	serie.ID = util.ParseID(pushed.URL())

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

// Delete a serie in Firebase
func (serie *Serie) Delete(log logging.Logger) (err error) {
	err = internal.Firebase.Child(pathSeries).Child(serie.ID).Remove()
	internal.LogRemove(log, pathSeries, serie.ID)
	return
}
