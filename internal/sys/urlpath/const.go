package urlpath

// Path pieces
const (
	Series   = "/series"
	Login    = "/login"
	Logout   = "/logout"
	Callback = "/callback"
)

// Path params
const (
	IDPathParam = "ID"
	IDPathDef   = "{" + IDPathParam + "}"
)

// Query params
const (
	CodeQueryParam  = "code"
	StateQueryParam = "state"
)

// Cookies
const (
	AuthChecksumCookieParameter = "auth_checksum"
)
