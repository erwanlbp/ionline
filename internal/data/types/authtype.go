package types

import (
	"strings"

	"github.com/erwanlbp/ionline/internal/sys/urlpath"
)

// AuthChecksum represents the cookie auth_checksum as type
type AuthChecksum string

// Cookie creates the cookie representation out of the AuthChecksum type
func (a AuthChecksum) Cookie() string {
	return urlpath.AuthChecksumCookieParameter + "=" + string(a)
}

// ExtractAuthChecksum extract the AuthChecksum out of a cookie if in.
func ExtractAuthChecksum(cookie string) (authChecksum AuthChecksum, ok bool) {
	if strings.HasPrefix(cookie, urlpath.AuthChecksumCookieParameter) {
		s := strings.TrimPrefix(cookie, urlpath.AuthChecksumCookieParameter+"=")
		s = strings.Split(s, ";")[0]
		authChecksum = AuthChecksum(s)
		ok = true
	}
	return
}
