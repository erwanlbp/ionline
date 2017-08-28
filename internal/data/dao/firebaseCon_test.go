package dao_test

import (
	"testing"

	"gopkg.in/zabawaba99/firego.v1"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/data/dao/internal"
	"github.com/erwanlbp/ionline/internal/test"
)

func TestSetFirebaseClient(t *testing.T) {
	assert := test.InitInternal(t)

	fbtest := firego.New("http://firebase.test", nil)
	oldFB := dao.SetFirebaseClient(fbtest)
	defer dao.SetFirebaseClient(oldFB)
	assert.Equal(internal.Firebase.URL(), "http://firebase.test")
}
