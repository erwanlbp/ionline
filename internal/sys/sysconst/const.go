package sysconst

import "time"

// Various constants
const (
	SiteName = "IOnline"

	// Authentication
	AuthChecksumExpire = time.Hour * 24
	AuthCookieName     = "auth_checksum"
)
