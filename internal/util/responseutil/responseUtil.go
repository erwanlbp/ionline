package responseutil

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
)

type returnType int

const (
	nothing = returnType(iota)
	dto
	templater
	redirect
	file
	err
)

// ReturnData represent the type of the data returned to the client and all the fields needed for every type
type ReturnData struct {
	typer     returnType
	path      string
	cookie    *http.Cookie
	data      interface{}
	errorMsg  error
	errorCode int
}

// Nothing will return a code 200
func Nothing() *ReturnData {
	return &ReturnData{
		typer: nothing,
	}
}

// Dto will return a JSON object with a code 204
func Dto(data interface{}) *ReturnData {
	return &ReturnData{
		typer: dto,
		data:  data,
	}
}

// Template will return a template from a html file, with a code 200
func Template(path string, data interface{}) *ReturnData {
	return &ReturnData{
		typer: templater,
		path:  path,
		data:  data,
	}
}

// Redirect will redirect to the path with a code 303
func Redirect(path string) *ReturnData {
	return &ReturnData{
		typer: redirect,
		path:  path,
	}
}

// File will serve the given filepath with a code 200
func File(path string) *ReturnData {
	return &ReturnData{
		typer: file,
		path:  path,
	}
}

// Error will return an error with the given code and error
func Error(code int, message error) *ReturnData {
	return &ReturnData{
		typer:     err,
		errorMsg:  message,
		errorCode: code,
	}
}

// Error will return the error if there is one
func (r *ReturnData) Error() error {
	return r.errorMsg
}

// AddCookie will add a cookie to the return datas
func (r *ReturnData) AddCookie(cookie *http.Cookie) *ReturnData {
	cookie.Path = "/"
	r.cookie = cookie
	return r
}

// SendResponse writes the response and the possible error in the HTTP response writer
func SendResponse(returnData *ReturnData, log logging.ExtendedLogger, writer http.ResponseWriter, request *http.Request) {
	var responseCode int

	// ReturnData shouldn't be nil
	if returnData == nil {
		log.Warnln("ReturnData shouldn't be nil")
		returnData = Nothing()
		return
	}

	// Add cookie if needed
	if returnData.cookie != nil {
		http.SetCookie(writer, returnData.cookie)
	}

	switch returnData.typer {
	case nothing:
		writer.WriteHeader(http.StatusNoContent)
		responseCode = http.StatusNoContent

	case dto:
		js, err := json.Marshal(returnData.data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			responseCode = http.StatusInternalServerError
			logError(log, http.StatusInternalServerError, err.Error())
			break
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(js)
		responseCode = http.StatusOK

	case templater:
		tmpl := template.Must(template.ParseFiles(returnData.path))
		tmpl.Execute(writer, returnData.data)
		responseCode = http.StatusOK

	case file:
		http.ServeFile(writer, request, returnData.path)
		responseCode = http.StatusOK

	case err:
		if returnData.errorCode == 0 {
			log.Warnln("ErrorCode should be set (errorMsg=" + returnData.errorMsg.Error() + ")")
			returnData.errorCode = http.StatusBadRequest
		}
		http.Error(writer, returnData.errorMsg.Error(), returnData.errorCode)
		responseCode = returnData.errorCode
		logError(log, returnData.errorCode, returnData.errorMsg.Error())

	case redirect:
		responseCode = http.StatusTemporaryRedirect
		if returnData.path == "" {
			log.Warnln("Redirect path shouldn't be \"\"")
			returnData.path = urlpath.IndexClientURL()
		}
		http.Redirect(writer, request, returnData.path, responseCode)

	}

	log.Println("Request status:", responseCode, http.StatusText(responseCode))

	return
}

func logError(log logging.ExtendedLogger, code int, msg string) {
	if code >= 500 {
		log.Critical(msg)
	} else if code >= 400 {
		log.Error(msg)
	}
}
