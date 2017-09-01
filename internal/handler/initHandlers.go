package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/sys/sysauth"
	"github.com/erwanlbp/ionline/internal/sys/sysconst"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/util/argutil"
	"github.com/erwanlbp/ionline/internal/util/responseutil"
)

// Init initializes the REST routes for the website
func Init() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	initIndexPage(router)
	initSeries(router)
	initLoginPage(router)

	return router
}

// InitRequestEnvironment is a wrapper around a Http.HandlerFunc
// it add a Logger and take care of sending the response and the error
func InitRequestEnvironment(action func(logging.Logger, *argutil.Args) *responseutil.ReturnData) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		args := argutil.New(request)

		log := logging.NewLogger()
		log.Println(request.Method, ":", request.RequestURI)

		returnData := action(log, &args)
		responseutil.SendResponse(returnData, log, writer, request)

		elapsed := time.Since(start)
		log.Println("Request time:", elapsed)

		go func() {
			err := log.Release()
			if err != nil {
				fmt.Println("ERROR: ", err)
			}
		}()
	}
}

// RequireAuthentify get the auth cookie and check that he is still valid
func RequireAuthentify(action func(logging.Logger, *argutil.Args) *responseutil.ReturnData) func(logging.Logger, *argutil.Args) *responseutil.ReturnData {
	return func(log logging.Logger, args *argutil.Args) *responseutil.ReturnData {
		cookieAuth := args.Cookie(sysconst.AuthCookieName)
		if args.Error() != nil {
			log.Println("Cookie auth_state not found => redirecting to", urlpath.LoginPageClientURL())
			return responseutil.Redirect(urlpath.LoginPageClientURL())
		}

		email, ok := sysauth.IsAuthentified(cookieAuth.Value)
		if !ok {
			log.Println("Unauthentified => redirecting to login")
			return responseutil.Redirect(urlpath.LoginPageClientURL())
		}

		log.Println("User authentified:", email)
		return action(log, args)
	}
}
