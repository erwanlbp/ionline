package host

import (
	"errors"
	"strings"
)

// Host represents a supported website where to find series
type Host int

// Supported hosts
const (
	ZoneTelechargementWs = Host(iota)
)

// Strings for the supported hosts
const (
	ZoneTelechargementWsString = "zone-telechargement.ws"
)

// HostName maps Host -> String
var HostName = map[Host]string{
	ZoneTelechargementWs: ZoneTelechargementWsString,
}

// Parser offers functions to find infos on the pages of the host
type Parser interface {
	Name() string
	Host() Host
	LastEpisode(string) (int, error)
	Season(string) (int, error)
}

// GetParser returns the GetParser matching the url's host
func GetParser(url string) (host Parser, err error) {
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "www.")
	urlSplit := strings.Split(url, "/")

	switch urlSplit[0] {
	case ZoneTelechargementWsString:
		host = &ZoneTelechargementWsParser{}
	default:
		err = errors.New("Unknown host " + urlSplit[0])
	}
	return
}
