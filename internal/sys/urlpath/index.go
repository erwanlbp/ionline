package urlpath

// IndexPath returns the path to the index with the possible params
func IndexPath() string {
	return indexBasePath()
}

// IndexClientURL returns the URL of the index with the params replaced
func IndexClientURL() string {
	return indexBasePath()
}

func indexBasePath() string {
	return "/"
}
