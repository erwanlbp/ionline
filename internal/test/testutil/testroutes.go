package testutil

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/erwanlbp/ionline/internal/handler"
	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/util/argutil"
	"github.com/erwanlbp/ionline/internal/util/responseutil"
)

const (
	// TestRouteClientURL is the path of the test route
	TestRouteClientURL = "/test-route"
)

// InitTestRoutes create the test routes
func InitTestRoutes(router *mux.Router) {
	router.Methods(http.MethodGet).Path(TestRouteClientURL).HandlerFunc(handler.InitRequestEnvironment(handler.RequireAuthentify(basicRoute)))
}

func basicRoute(log logging.Logger, args *argutil.Args) *responseutil.ReturnData {
	return responseutil.Nothing()
}
