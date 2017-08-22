package sys

import (
	"os"
	"path/filepath"
)

func init() {
	// If we are not located in the right directory we stop the server
	path, err := filepath.Abs(PublicPath())
	if err == nil {
		_, err = os.Stat(path)
		os.IsNotExist(err)
	}

	if err != nil {
		panic(err.Error())
	}
}

// PublicPath returns the path to the html templates
func PublicPath() string {
	return rootInternal() + "public/www/"
}

func rootInternal() string {
	return "internal/"
}
