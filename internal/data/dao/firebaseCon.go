package dao

import (
	"gopkg.in/zabawaba99/firego.v1"

	"github.com/erwanlbp/ionline/internal/data/dao/internal"
)

// SetFirebaseClient fill the Firebase variable with the Firebase object received
func SetFirebaseClient(firebase *firego.Firebase) {
	internal.Firebase = firebase
}
