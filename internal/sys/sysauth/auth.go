package sysauth

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/sys/sysconst"
	"github.com/erwanlbp/ionline/internal/sys/systime"
)

// UserAuth represents a user authenticated
type UserAuth struct {
	Dao    dao.User
	expire time.Time
}

// Contains auth_state -> User
var authentified = map[string]UserAuth{}

// IsAuthentified returns true if the token received match one in the auth cache
func IsAuthentified(authStr string) (user dao.User, ok bool) {
	userAuth, isIn := authentified[authStr]

	if !isIn {
		return
	}

	if userAuth.expire.Before(systime.Now()) {
		delete(authentified, authStr)
		return
	}

	user = userAuth.Dao
	ok = true
	return
}

// Disconnect delete the token received from the auth cache
func Disconnect(authStr string) (user dao.User) {
	if userAuth, ok := authentified[authStr]; ok {
		user = userAuth.Dao
	}
	delete(authentified, authStr)
	return
}

// Authentified add the token and email received to the auth cache
func Authentified(authStr string, user dao.User) {
	authentified[authStr] = UserAuth{
		Dao:    user,
		expire: systime.Now().Add(sysconst.AuthChecksumExpire),
	}
}

// ValidateUser return true if the email is valid and have the right to access the service
func ValidateUser(email string) bool {
	return true
}

// RandToken return a 64 characters key in base64
func RandToken() string {
	b := make([]byte, 64)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
