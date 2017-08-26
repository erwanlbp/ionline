package extdep

import "github.com/maprost/restclient"

type serie interface {
	Get(string) (string, error)
}

// Serie is the external dependency for the series
// It offers some function to retrieve data from the web
var Serie serie = &serieImpl{}

type serieImpl struct{}

// Get the content from the page with that URL
func (s *serieImpl) Get(url string) (string, error) {
	pageContent, result := restclient.Get(url).NoLogger().SendAndGetResponse()
	return pageContent, result.Err
}
