package host

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ZoneTelechargementWsParser is the parser for zone-telechargement.ws
type ZoneTelechargementWsParser struct {
}

// Name of the host : zone-telechargement.ws
func (h *ZoneTelechargementWsParser) Name() string {
	return ZoneTelechargementWsString
}

// Host (int)
func (h *ZoneTelechargementWsParser) Host() Host {
	return ZoneTelechargementWs
}

// LastEpisode finds the episode by looking for the list of all episodes per host (e.g Uptobox, ...)
func (h *ZoneTelechargementWsParser) LastEpisode(page string) (lastEpisode int, err error) {
	regex := regexp.MustCompile(">Episode [0-9]+(| FiNAL| Final)<")
	matches := regex.FindAllString(page, -1)
	if len(matches) == 0 {
		err = errors.New("Can't find an episode on the page")
		return
	}

	match := matches[len(matches)-1]

	matchSplit := strings.Split(match, " ")
	matchSplit[1] = strings.TrimSuffix(matchSplit[1], "<")
	episode, err := strconv.ParseInt(matchSplit[1], 10, 8)
	if err != nil {
		err = errors.New("Can't find the episode in the last episode occurence : " + err.Error())
		return
	}
	lastEpisode = int(episode)

	return
}

// Season finds the seaon on the page in the <title> tag
func (h *ZoneTelechargementWsParser) Season(page string) (season int, err error) {
	// Find the <title>
	regexTitle := regexp.MustCompile("<title>[^<]+</title>")
	title := regexTitle.FindString(page)
	if title == "" {
		err = errors.New("Can't find the title tag for the season")
		return
	}

	// Find the season in the title
	regex := regexp.MustCompile("Saison [0-9]+")
	match := regex.FindString(title)
	if match == "" {
		err = errors.New("Can't find the season in title " + title)
		return
	}

	// Find the season number
	matchSplit := strings.Split(match, " ")
	season64, err := strconv.ParseInt(matchSplit[1], 10, 8)
	if err != nil {
		err = errors.New("Finding season: " + err.Error())
		return
	}
	season = int(season64)

	return
}
