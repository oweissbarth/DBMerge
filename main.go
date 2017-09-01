package main

import (
	"github.com/op/go-logging"
	"github.com/oweissbarth/DBMerge/view"
)

var log = logging.MustGetLogger("DBMergeMain")

func main() {
	log.Info("Starting DBMerge")
	view.Construct()
}
