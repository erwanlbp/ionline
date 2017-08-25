package internal

import (
	"strings"

	"gopkg.in/zabawaba99/firego.v1"

	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/util"
)

// Firebase instance
var Firebase *firego.Firebase

// LogPush add a log describing the action
func LogPush(log logging.Logger, fb *firego.Firebase, v interface{}) {
	log.Println("Push to", util.ParsePath(fb.String()), v)
}

// LogValue add a log describing the action
func LogValue(log logging.Logger, path ...string) {
	log.Println("Read value at", strings.Join(path, "/")+"/")
}

// LogRemove add a log describing the action
func LogRemove(log logging.Logger, path ...string) {
	log.Println("Delete value at", strings.Join(path, "/"))
}
