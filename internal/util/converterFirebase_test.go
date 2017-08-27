package util_test

import (
	"testing"

	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/util"
)

func TestParseID(t *testing.T) {
	assert := test.InitInternal(t)

	testdatas := []struct {
		url string
		id  string
	}{
		{"http://ionline-test.firebase.io/series/id-of-the-serie.json", "id-of-the-serie"},
		{"http://ionline-test.firebase.io/id-of-the-serie.json", "id-of-the-serie"},
		{"http://ionline-test.firebase.io/id-of-the-serie", "id-of-the-serie"},
	}

	for _, testdata := range testdatas {
		assert.Equal(util.ParseID(testdata.url), testdata.id)
	}
}
