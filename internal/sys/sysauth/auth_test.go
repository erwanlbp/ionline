package sysauth_test

import (
	"testing"
	"time"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/sys/sysauth"
	"github.com/erwanlbp/ionline/internal/sys/sysconst"
	"github.com/erwanlbp/ionline/internal/sys/systime"
	"github.com/erwanlbp/ionline/internal/test"
	"github.com/erwanlbp/ionline/internal/test/testmock"
)

func TestValidateUser(t *testing.T) {
	assert := test.InitInternal(t)

	var datas = []struct {
		email string
		valid bool
	}{
		{"erwan.lbp@gmail.com", true},
	}

	for _, data := range datas {
		assert.Equal(sysauth.ValidateUser(data.email), data.valid)
	}
}

func TestAuthentification(t *testing.T) {
	assert := test.InitInternal(t)

	randKey := sysauth.RandToken()

	_, ok := sysauth.IsAuthentified(randKey)
	assert.False(ok)

	user := dao.User{Email: testmock.EmailTest}

	sysauth.Authentified(randKey, user)

	userAuth, ok := sysauth.IsAuthentified(randKey)
	assert.True(ok)
	assert.Equal(userAuth, user)
}

func TestIsAuthentified_expired(t *testing.T) {
	assert := test.InitInternal(t)

	systime.InitTime(time.Now())

	randKey := sysauth.RandToken()
	sysauth.Authentified(randKey, dao.User{Email: testmock.EmailTest})

	_, ok := sysauth.IsAuthentified(randKey)
	assert.True(ok)

	systime.InitTime(time.Now().Add(sysconst.AuthChecksumExpire))

	_, ok = sysauth.IsAuthentified(randKey)
	assert.False(ok)
}

func TestAuthentified_BadKey(t *testing.T) {
	assert := test.InitInternal(t)

	randKey := sysauth.RandToken()

	sysauth.Authentified(randKey, dao.User{Email: testmock.EmailTest})

	_, ok := sysauth.IsAuthentified(randKey + "Bad")
	assert.False(ok)
}

func TestDisconnect(t *testing.T) {
	assert := test.InitInternal(t)

	randKey := sysauth.RandToken()

	user := dao.User{Email: testmock.EmailTest}

	sysauth.Authentified(randKey, user)
	_, ok := sysauth.IsAuthentified(randKey)
	assert.True(ok)

	userdisconnect := sysauth.Disconnect(randKey)
	assert.Equal(userdisconnect, user)
	_, ok = sysauth.IsAuthentified(randKey)
	assert.False(ok)
}
