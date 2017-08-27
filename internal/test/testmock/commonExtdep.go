package testmock

import (
	"fmt"
	"net/url"

	"github.com/erwanlbp/ionline/internal/data/host"
)

// MockedCommonExtdep mocks the extdep.CommonClient
type MockedCommonExtdep struct {
}

// Get the mocked content of the page
// It parse the URL to find which content to return
func (m *MockedCommonExtdep) Get(pageURL string) (pageContent string, err error) {
	fmt.Println("[mock] Get " + pageURL)

	urlParsed, err := url.Parse(pageURL)
	if err != nil {
		return
	}

	domain := urlParsed.Hostname()
	switch domain {
	case host.ZoneTelechargementWsString:
		pageContent = "<html>\n<head>\n" +
			"<title>Telecharger Game of Thrones - Saison 5 gratuit Zone Telechargement - Site de Téléchargement Gratuit</title>\n" +
			"<meta name='description' content='Telecharger Game of ThronesQualité HD 720p | FRENCH  Saison 5 Complete    Origine de la serie :  Américaine Réalisateur :  D.B. Weiss Acteurs :  Peter Dinklage, Nikolaj Coster-Waldau, Lena Headey Genre'/>\n" +
			"</head>\n<body>\n" +
			"<font color=red>Game.of.Thrones.S05.FRENCH.720p.WEB-DL.DD5.1.H264-SVR</font><br /><b><div>Uptobox</div></b><b><a href='https://www.dl-protecte.org/1234556001234556021234556101234556157k4ykpaplz6s\n" +
			">Episode 1</a></b><br /><b><a href='https://www.dl-protecte.org/123455600123455602123455610123455615kd8qz2b3aisz'\n" +
			">Episode 10 Final</a></b><br /><br /><b><div>Uploaded</div></b><b><a href='https://www.dl-protecte.org/123455600123455605123455615azbducyw'\n" +
			"</body>\n</head>\n</html>"
	default:
		pageContent = "<html>\n<head>\n<title>some title</title>\n</head>\n<body>The body</body>\n</html>"
	}

	// Exceptions for the error cases
	switch urlParsed.Path {
	case "/zt-wrong-season":
		pageContent = "<html>\n<head><title>something not the season</title></head>\n<body>\n" +
			"<font color=red>Game.of.Thrones.S05.FRENCH.720p.WEB-DL.DD5.1.H264-SVR</font><br /><b><div>Uptobox</div></b><b><a href='https://www.dl-protecte.org/1234556001234556021234556101234556157k4ykpaplz6s\n" +
			">Episode 1</a></b><br /><b><a href='https://www.dl-protecte.org/123455600123455602123455610123455615kd8qz2b3aisz'\n" +
			">Episode 10 Final</a></b><br /><br /><b><div>Uploaded</div></b><b><a href='https://www.dl-protecte.org/123455600123455605123455615azbducyw'\n" +
			"</body>\n</head>\n</html>"
	case "/zt-wrong-episode":
		pageContent = "<html>\n<head>\n" +
			"<title>Telecharger Game of Thrones - Saison 5 gratuit Zone Telechargement - Site de Téléchargement Gratuit</title>\n" +
			"<meta name='description' content='Telecharger Game of ThronesQualité HD 720p | FRENCH  Saison 5 Complete    Origine de la serie :  Américaine Réalisateur :  D.B. Weiss Acteurs :  Peter Dinklage, Nikolaj Coster-Waldau, Lena Headey Genre'/>\n" +
			"</head>\n<body>\n" +
			"<font color=red>Game.of.Thrones.S05.FRENCH.720p.WEB-DL.DD5.1.H264-SVR</font><br /><b><div>Uptobox</div></b><b><a href='https://www.dl-protecte.org/1234556001234556021234556101234556157k4ykpaplz6s\n" +
			"</body>\n</head>\n</html>"
	}

	return
}
