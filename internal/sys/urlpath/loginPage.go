package urlpath

// ============== LoginPage ===================

// LoginPageBoundaryPath returns the string of the base path to access the login page
func LoginPageBoundaryPath() string {
	return loginPageBasePath()
}

// LoginPageClientURL returns the string of the URL to access the login page
func LoginPageClientURL() string {
	return loginPageBasePath()
}

func loginPageBasePath() string {
	return Login
}

// ============== Logout ===================

// LogoutBoundaryPath returns the string of the base path to access the logout
func LogoutBoundaryPath() string {
	return logoutBasePath()
}

// LogoutClientURL returns the string of the URL to access the logout
func LogoutClientURL() string {
	return logoutBasePath()
}

func logoutBasePath() string {
	return Logout
}

// ============== Auth ===================

// AuthBoundaryPath returns the string of the base path to access the Auth request
func AuthBoundaryPath() string {
	return authBasePath()
}

// AuthClientURL returns the string of the URL to access the Auth request
func AuthClientURL() string {
	return authBasePath()
}

func authBasePath() string {
	return Login + Callback
}
