package responseutil

import (
	"github.com/erwanlbp/ionline/internal/sys/sysconst"
	"github.com/erwanlbp/ionline/internal/sys/urlpath"
)

// BaseTemplateData represents the basic data that are needed by every page, like the SiteName
type BaseTemplateData struct {
	SiteName string
	Header   string
	IndexURL string
}

// BaseTemplateDatas returns the object initialized with the base infos
// To complete it, use the FillXXX functions
func BaseTemplateDatas() BaseTemplateData {
	return BaseTemplateData{
		SiteName: sysconst.SiteName,
		Header:   sysconst.SiteName,
		IndexURL: urlpath.IndexClientURL(),
	}
}

// FillHeader add the header to the datas
func (btd BaseTemplateData) FillHeader(header string) BaseTemplateData {
	btd.Header = btd.SiteName + " - " + header
	return btd
}
