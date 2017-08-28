package util

import "strings"

// ParseID takes a Firebase URL and return just the last part, which is the ID
func ParseID(url string) (id string) {
	urlSplit := strings.Split(url, "/")
	id = urlSplit[len(urlSplit)-1]
	id = strings.TrimSuffix(id, ".json")
	return
}
