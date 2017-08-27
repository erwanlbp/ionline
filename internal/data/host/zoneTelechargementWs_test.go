package host_test

import (
	"testing"

	"github.com/erwanlbp/ionline/internal/data/host"
	"github.com/erwanlbp/ionline/internal/test"
)

func TestZoneTelechargementWsParser_Host(t *testing.T) {
	assert := test.InitInternal(t)

	parser := host.ZoneTelechargementWsParser{}
	assert.Equal(parser.Host(), host.ZoneTelechargementWs)
}

func TestZoneTelechargementWsParser_Name(t *testing.T) {
	assert := test.InitInternal(t)

	parser := host.ZoneTelechargementWsParser{}
	assert.Equal(parser.Name(), "www.zone-telechargement.ws")
}

func TestZoneTelechargementWsParser_LastEpisode(t *testing.T) {
	assert := test.InitInternal(t)

	parser := host.ZoneTelechargementWsParser{}

	// An extract from a page
	pageContent := "<html>\n<head>\n" +
		"<title>Telecharger Game of Thrones - Saison 5 gratuit Zone Telechargement - Site de Téléchargement Gratuit</title>\n" +
		"<meta name='description' content='Telecharger Game of ThronesQualité HD 720p | FRENCH  Saison 5 Complete    Origine de la serie :  Américaine Réalisateur :  D.B. Weiss Acteurs :  Peter Dinklage, Nikolaj Coster-Waldau, Lena Headey Genre'/>\n" +
		"</head>\n<body>\n" +
		"<font color=red>Game.of.Thrones.S05.FRENCH.720p.WEB-DL.DD5.1.H264-SVR</font><br /><b><div>Uptobox</div></b><b><a href='https://www.dl-protecte.org/1234556001234556021234556101234556157k4ykpaplz6s\n" +
		">Episode 1</a></b><br /><b><a href='https://www.dl-protecte.org/123455600123455602123455610123455615kd8qz2b3aisz'\n" +
		">Episode 10 Final</a></b><br /><br /><b><div>Uploaded</div></b><b><a href='https://www.dl-protecte.org/123455600123455605123455615azbducyw'\n" +
		"</body>\n</head>\n</html>"

	lastEpisode, err := parser.LastEpisode(pageContent)
	assert.Nil(err)
	assert.Equal(lastEpisode, 10)
}

func TestZoneTelechargementWsParser_LastEpisode_NotFound(t *testing.T) {
	assert := test.InitInternal(t)

	parser := host.ZoneTelechargementWsParser{}
	pageContent := "<html>\n<head><title>something not the season</title></head>\n" +
		"<body>Something not the episode</body></html>"

	_, err := parser.LastEpisode(pageContent)
	assert.NotNil(err)
}

func TestZoneTelechargementWsParser_Season(t *testing.T) {
	assert := test.InitInternal(t)

	parser := host.ZoneTelechargementWsParser{}

	// An extract from a page
	pageContent := "<html>\n<head>\n" +
		"<title>Telecharger Game of Thrones - Saison 5 gratuit Zone Telechargement - Site de Téléchargement Gratuit</title>\n" +
		"<meta name='description' content='Telecharger Game of ThronesQualité HD 720p | FRENCH  Saison 5 Complete    Origine de la serie :  Américaine Réalisateur :  D.B. Weiss Acteurs :  Peter Dinklage, Nikolaj Coster-Waldau, Lena Headey Genre'/>\n" +
		"</head>\n<body>\n" +
		"<font color=red>Game.of.Thrones.S05.FRENCH.720p.WEB-DL.DD5.1.H264-SVR</font><br /><b><div>Uptobox</div></b><b><a href='https://www.dl-protecte.org/1234556001234556021234556101234556157k4ykpaplz6s\n" +
		">Episode 1</a></b><br /><b><a href='https://www.dl-protecte.org/123455600123455602123455610123455615kd8qz2b3aisz'\n" +
		">Episode 10 Final</a></b><br /><br /><b><div>Uploaded</div></b><b><a href='https://www.dl-protecte.org/123455600123455605123455615azbducyw'\n" +
		"</body>\n</head>\n</html>"

	season, err := parser.Season(pageContent)
	assert.Nil(err)
	assert.Equal(season, 5)
}

func TestZoneTelechargementWsParser_Season_NotFound(t *testing.T) {
	assert := test.InitInternal(t)

	parser := host.ZoneTelechargementWsParser{}
	pageContent := "<html>\n<head><title>something not the season</title></head><body>Something still not the season</body></html>"

	_, err := parser.Season(pageContent)
	assert.NotNil(err)
}
