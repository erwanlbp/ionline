package handler_test

import (
	"testing"

	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/test/testrest"
)

func TestGetIndex(t *testing.T) {
	assert, _ := test.InitRest(t)

	testrest.GetIndex200(assert)
}
