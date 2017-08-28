package webcrawler

import (
	"fmt"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/data/host"
	"github.com/erwanlbp/ionline/internal/extdep"
	"github.com/erwanlbp/ionline/internal/sys/logging"
)

// ActualizeSerie fetch the page content and parse it to find the infos needed
func ActualizeSerie(log logging.Logger, serie *dao.Serie) (err error) {
	// Find the right parser for the serie.URL
	hostParser, err := host.GetParser(serie.URL)
	if err != nil {
		return
	}
	serie.Host = hostParser.Host()
	log.Println("Host is", hostParser.Name())

	// Get the content of the serie.URL
	pageContent, err := extdep.CommonClient.Get(serie.URL)
	if err != nil {
		return
	}

	// Find the season
	season, err := hostParser.Season(pageContent)
	if err != nil {
		return
	}
	serie.Season = season
	log.Println(fmt.Sprintf("Season is %v", season))

	// Find the last episode out
	episode, err := hostParser.LastEpisode(pageContent)
	if err != nil {
		return
	}
	serie.LastEpisode = episode
	log.Println(fmt.Sprintf("Last episode is %v", episode))

	return
}
