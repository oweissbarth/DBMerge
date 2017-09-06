package main

import (
	internalLog "log"

	"github.com/op/go-logging"
	"github.com/oweissbarth/DBMerge/view"
)

var log = logging.MustGetLogger("DBMergeMain")

func main() {
	internalLog.SetFlags(internalLog.LstdFlags | internalLog.Lshortfile)

	log.Info("Starting DBMerge")
	view.Construct()
}
