package util

import "strings"

// ParseID takes a Firebase URL and return just the last part, which is the ID
func ParseID(url string) (id string) {
	urlSplit := strings.Split(url, "/")
	id = urlSplit[len(urlSplit)-1]
	return
}

// ParsePath takes a Firebase URL and remove the host and the .json at the end
func ParsePath(url string) (path string) {
	sep := ".com/"
	endHost := strings.Index(url, sep)
	path = url[endHost+len(sep):]
	path = strings.TrimSuffix(path, ".json")
	return
}
