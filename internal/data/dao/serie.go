package dao

import (
	"github.com/erwanlbp/ionline/internal/data/dao/internal"
	"github.com/erwanlbp/ionline/internal/data/protobuf"
	"github.com/erwanlbp/ionline/internal/sys/logging"
	"github.com/erwanlbp/ionline/internal/util"
)

// AddSerie push a serie in the Firebase path /series
func AddSerie(log logging.Logger, serie *protobuf.Serie) (err error) {
	pushed, err := internal.Firebase.Child("series").Push(serie)
	internal.LogPush(log, pushed, serie)
	if err != nil {
		return
	}

	serie.Id = util.ParseID(pushed.URL())

	return
}
