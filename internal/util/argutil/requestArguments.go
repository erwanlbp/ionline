package argutil

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

// Args contains the request and the errors found during the parse
// It provides functions to retrieve the arguments of the request
type Args struct {
	request *http.Request
	err     error
}

// New returns a Args object containing the request
func New(request *http.Request) Args {
	return Args{
		request: request,
	}
}

func (ra *Args) addError(err error) {
	if err != nil && ra.err == nil {
		ra.err = err
	}
}

// StringPathParam returns the request argument with the pathParam name
func (ra *Args) StringPathParam(name string) string {
	arg, isIn := mux.Vars(ra.request)[name]
	if !isIn {
		ra.addError(errors.New("Path param " + name + " not found"))
		return ""
	}
	arg, err := url.QueryUnescape(arg)
	ra.addError(err)
	return arg
}

// StringQueryParamOptional returns the request url parameter matching name
func (ra *Args) StringQueryParamOptional(name string) (string, bool) {
	arg := ra.request.URL.Query().Get(name)
	return arg, arg != ""
}

// Int64PathParam returns the request argument with the pathParam name
func (ra *Args) Int64PathParam(name string) int64 {
	argStr := ra.StringPathParam(name)
	argInt, err := strconv.ParseInt(argStr, 10, 64)
	ra.addError(err)
	return argInt
}

// StringQueryParam returns the request url parameter matching name
func (ra *Args) StringQueryParam(name string) string {
	arg, isIn := ra.StringQueryParamOptional(name)
	if !isIn {
		ra.addError(errors.New("Query parameter " + name + " not found"))
	}
	return arg
}

// Int64QueryParamOptional returns the int64 request url parameter matching name
func (ra *Args) Int64QueryParamOptional(name string) (int64, bool) {
	argStr, isIn := ra.StringQueryParamOptional(name)
	arg, err := strconv.ParseInt(argStr, 10, 64)
	return arg, isIn && err == nil
}

// HeaderOptional return the request header name
func (ra *Args) HeaderOptional(name string) (string, bool) {
	header := ra.request.Header.Get(name)
	header, err := url.QueryUnescape(header)
	return header, (err == nil && header != "")
}

// Header return the request header name
func (ra *Args) Header(name string) string {
	header, isIn := ra.HeaderOptional(name)
	if !isIn {
		ra.addError(errors.New("Header " + name + " not found"))
	}
	return header
}

// Body returns the body of the request
func (ra *Args) Body(decoded interface{}) {
	err := json.NewDecoder(ra.request.Body).Decode(&decoded)
	ra.addError(err)
}

// Cookie return the cookie with the given name or nil if not found
func (ra *Args) Cookie(name string) *http.Cookie {
	cookie, errNotFound := ra.request.Cookie(name)
	ra.addError(errNotFound)
	return cookie
}

// Errors returns the errors that were found
func (ra *Args) Error() error {
	return ra.err
}
