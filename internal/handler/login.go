package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/erwanlbp/ionline/internal/extdep"
	"github.com/erwanlbp/ionline/internal/sys"
	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/sys/sysauth"
	"github.com/erwanlbp/ionline/internal/sys/sysconst"
	"github.com/erwanlbp/ionline/internal/sys/systime"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/util/argutil"
	"github.com/erwanlbp/ionline/internal/util/responseutil"
)

// LoginPageData represents the data to fill the template of the login page
type LoginPageData struct {
	BaseData responseutil.BaseTemplateData
	LoginURL string
}

func initLoginPage(router *mux.Router) {
	router.Methods(http.MethodGet).Path(urlpath.LoginPageClientURL()).HandlerFunc(InitRequestEnvironment(loginPage))
	router.Methods(http.MethodGet).Path(urlpath.AuthClientURL()).HandlerFunc(InitRequestEnvironment(authHandler))
	router.Methods(http.MethodGet).Path(urlpath.LogoutClientURL()).HandlerFunc(InitRequestEnvironment(RequireAuthentify(logout)))
}

func loginPage(log logging.Logger, args *argutil.Args) *responseutil.ReturnData {
	cookieAuth := args.Cookie(urlpath.AuthChecksumCookieParameter)

	// If user is logged in, he can't access this page
	if args.Error() == nil {
		email, ok := sysauth.IsAuthentified(cookieAuth.Value)
		if ok {
			log.Println("User", email, "is already authentified")
			return responseutil.Redirect(urlpath.IndexClientURL())
		}
	}

	state := sysauth.RandToken()
	loginURL := extdep.GoogleAuthClient.GetLoginURL(state)

	return responseutil.Template(sys.PagePath()+"login.html",
		LoginPageData{
			BaseData: responseutil.BaseTemplateDatas().FillHeader("Login"),
			LoginURL: loginURL,
		}).
		AddCookie(&http.Cookie{
			Name:    urlpath.AuthChecksumCookieParameter,
			Value:   state,
			Expires: systime.Now().Add(sysconst.AuthChecksumExpire),
		})
}

func authHandler(log logging.Logger, args *argutil.Args) *responseutil.ReturnData {

	cookieState := args.Cookie(urlpath.AuthChecksumCookieParameter)
	queryState := args.StringQueryParam(urlpath.StateQueryParam)
	queryCode := args.StringQueryParam(urlpath.CodeQueryParam)

	if err := args.Error(); err != nil || cookieState.Value != queryState {
		log.Println("Cookie state and query state are not equals, or code is missing.", err.Error())
		return responseutil.Redirect(urlpath.LoginPageClientURL())
	}

	// Get user infos from Google
	user, err := extdep.GoogleAuthClient.GoogleAuthenticate(queryCode)
	if err != nil {
		log.Warnln("Problem during Google Authentification", err.Error())
		return responseutil.Redirect(urlpath.LoginPageClientURL())
	}

	// Check if the user has the right to connect to the service
	ok := sysauth.ValidateUser(user.Email)
	if !ok {
		log.Warnln("User email", user.Email, "isn't allowed to use the service")
		return responseutil.Redirect(urlpath.LoginPageClientURL())
	}

	// Authentify the user in the cache
	sysauth.Authentified(cookieState.Value, user)
	log.Println("User validated:", user.Email)

	return responseutil.Redirect(urlpath.IndexClientURL())
}

func logout(log logging.Logger, args *argutil.Args) *responseutil.ReturnData {
	authCookie := args.Cookie(urlpath.AuthChecksumCookieParameter)
	if err := args.Error(); err != nil {
		// Unreachable cause this handler require authentication => require cookie
		log.Warnln("AuthCookie not available.", err.Error())

	} else {
		emailLoggedOut := sysauth.Disconnect(authCookie.Value)
		log.Println("Logged out:", emailLoggedOut)
	}

	return responseutil.Redirect(urlpath.LoginPageClientURL())
}
