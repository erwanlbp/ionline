package extdep

import "github.com/maprost/restclient"

// Common functions to retrieve data from the web
// Can be mocked
type Common interface {
	Get(string) (string, error)
}

// CommonClient is the external dependency for the Common request
// It offers some function to retrieve data from the web
var CommonClient Common = &CommonImpl{}

// CommonImpl is the basic implementation for the Common interface
type CommonImpl struct{}

// Get the content from the page with that URL
func (s *CommonImpl) Get(url string) (string, error) {
	pageContent, result := restclient.Get(url).NoLogger().SendAndGetResponse()
	return pageContent, result.Err
}
