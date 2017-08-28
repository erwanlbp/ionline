package host_test

import (
	"testing"

	"github.com/erwanlbp/ionline/internal/data/host"
	"github.com/erwanlbp/ionline/internal/test"
)

func TestGetParser(t *testing.T) {
	assert := test.InitInternal(t)

	testdatas := []struct {
		url  string
		host host.Host
	}{
		{"http://www.zone-telechargement.ws/test/url", host.ZoneTelechargementWs},
		{"https://www.zone-telechargement.ws/test/url", host.ZoneTelechargementWs},
		{"http://www.zone-telechargement.ws", host.ZoneTelechargementWs},
		{"http://www.zone-telechargement.ws/", host.ZoneTelechargementWs},
	}

	for _, testdata := range testdatas {
		parser, err := host.GetParser(testdata.url)
		assert.Nil(err)
		assert.Equal(parser.Host(), testdata.host)
	}
}

func TestGetParser_UnknownHost(t *testing.T) {
	assert, _ := test.Init(t)

	_, err := host.GetParser("http://unknown-host/")
	assert.NotNil(err)
}
