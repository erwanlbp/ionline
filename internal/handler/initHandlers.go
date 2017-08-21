package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/util/argutil"
	"github.com/erwanlbp/ionline/internal/util/responseutil"
)

// Init initializes the REST routes for the website
func Init() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	initIndexPage(router)

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
