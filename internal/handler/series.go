package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/erwanlbp/ionline/internal/data/dao"
	"github.com/erwanlbp/ionline/internal/data/types"
	"github.com/erwanlbp/ionline/internal/sys"
	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
	"github.com/erwanlbp/ionline/internal/util/argutil"
	"github.com/erwanlbp/ionline/internal/util/responseutil"
)

type seriesListTemplateData struct {
	BaseDatas   responseutil.BaseTemplateData
	Series      []dao.Serie
	AddSerieURL string
	Languages   map[types.Language]string
	Qualities   map[types.Quality]string
}

// DeleteSerieURL returns the url containing the id of the serie to delete
func (t seriesListTemplateData) DeleteSerieURL(id string) string {
	return urlpath.DeleteSerieClientURL(id)
}

func initSeries(router *mux.Router) {
	router.Methods(http.MethodGet).Path(urlpath.SeriesPath()).HandlerFunc(InitRequestEnvironment(listSeries))
	router.Methods(http.MethodPost).Path(urlpath.AddSeriePath()).HandlerFunc(InitRequestEnvironment(addSerie))
	router.Methods(http.MethodDelete).Path(urlpath.DeleteSeriePath()).HandlerFunc(InitRequestEnvironment(deleteSerie))
}

func listSeries(log logging.Logger, args *argutil.Args) *responseutil.ReturnData {

	series, err := dao.FindAllSeries(log)
	if err != nil {
		return responseutil.Error(http.StatusInternalServerError, err)
	}

	return responseutil.Template(sys.PagePath()+"series_list.html",
		seriesListTemplateData{
			BaseDatas:   responseutil.BaseTemplateDatas(),
			Series:      series,
			AddSerieURL: urlpath.AddSerieClientURL(),
			Qualities:   types.QualityName,
			Languages:   types.LanguageName,
		})
}

func addSerie(log logging.Logger, args *argutil.Args) *responseutil.ReturnData {
	// Read the serie in the request
	body := args.Body()
	if args.Error() != nil {
		return responseutil.Error(http.StatusBadRequest, args.Error())
	}

	var newSerie dao.Serie
	newSerie.ParseJSON(body)

	// Save the new serie
	err := newSerie.Push(log)
	if err != nil {
		return responseutil.Error(http.StatusInternalServerError, err)
	}

	return responseutil.Nothing()
}

func deleteSerie(log logging.Logger, args *argutil.Args) *responseutil.ReturnData {
	// Read the ID of the serie in the path
	id := args.StringPathParam(urlpath.IDPathParam)
	if args.Error() != nil {
		return responseutil.Error(http.StatusBadRequest, args.Error())
	}

	// Delete the serie
	serie := &dao.Serie{ID: id}
	err := serie.Delete(log)
	if err != nil {
		return responseutil.Error(http.StatusInternalServerError, err)
	}

	return responseutil.Nothing()
}