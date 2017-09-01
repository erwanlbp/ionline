package sys

import (
	"github.com/erwanlbp/ionline/internal/sys/config"
)

// PagePath returns the path to the html templates
func PagePath() string {
	return config.ProjectPath() + "internal/public/www/"
}
