package sysauth_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/sys/sysauth"
	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/test/testinject"
)

func TestInitFirebase(t *testing.T) {
	assert := test.InitInternal(t)

	// Make sure that Firebase is wrongly setted
	test.WantFirebaseError()

	// Init Firebase so now it should be able to make requests
	err := sysauth.InitFirebase()
	assert.Nil(err)

	// Has to create unikey and logger there cause we didn't want the tests to be initialized
	unikey := fmt.Sprintf("%v-%v", t.Name(), time.Now().Nanosecond())
	log := logging.NewLogger()

	// Try to push a serie, if shouldn't fail
	testinject.PushSerie(assert, unikey, log)
}
