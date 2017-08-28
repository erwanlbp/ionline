package dao

import (
	"gopkg.in/zabawaba99/firego.v1"

	"github.com/erwanlbp/ionline/internal/data/dao/internal"
)

// SetFirebaseClient fill the Firebase variable with the Firebase object received and return the previous one
func SetFirebaseClient(firebase *firego.Firebase) *firego.Firebase {
	oldFB := internal.Firebase
	internal.Firebase = firebase
	return oldFB
}
