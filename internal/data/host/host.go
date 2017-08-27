package host

import (
	"errors"
	"net/url"
)

// Host represents a supported website where to find series
type Host int

// Supported hosts
const (
	ZoneTelechargementWs = Host(iota)
)

// Strings for the supported hosts
const (
	ZoneTelechargementWsString = "www.zone-telechargement.ws"
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
func GetParser(pageURL string) (host Parser, err error) {
	urlParsed, err := url.Parse(pageURL)
	if err != nil {
		return
	}
	domain := urlParsed.Hostname()

	switch domain {
	case ZoneTelechargementWsString:
		host = &ZoneTelechargementWsParser{}
	default:
		err = errors.New("Unknown host " + domain)
	}
	return
}
