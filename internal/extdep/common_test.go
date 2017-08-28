package extdep_test

import (
	"testing"

	"github.com/erwanlbp/ionline/internal/extdep"
	"github.com/erwanlbp/ionline/internal/test"
)

func TestCommonImpl_Get(t *testing.T) {
	assert := test.InitInternal(t)

	_, err := extdep.CommonClient.Get("http://www.google.com")
	assert.Nil(err)
}

func TestCommonImpl_Get_Unexistent_URL(t *testing.T) {
	assert := test.InitInternal(t)

	_, err := extdep.CommonClient.Get("http://127.0.1.1:9999")
	assert.NotNil(err)
}
