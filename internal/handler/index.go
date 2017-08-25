package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/erwanlbp/ionline/internal/sys"
	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/util/argutil"
	"github.com/erwanlbp/ionline/internal/util/responseutil"
)

type indexTemplateData struct {
	BaseDatas responseutil.BaseTemplateData
}

func initIndexPage(router *mux.Router) {
	router.Methods(http.MethodGet).Path(urlpath.IndexPath()).HandlerFunc(InitRequestEnvironment(indexPage))
}

func indexPage(log logging.Logger, args *argutil.Args) *responseutil.ReturnData {
	return responseutil.Template(sys.PagePath()+"index.html",
		indexTemplateData{
			BaseDatas: responseutil.BaseTemplateDatas(),
		})
}
