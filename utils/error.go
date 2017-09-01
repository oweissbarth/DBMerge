package utils

import logging "github.com/op/go-logging"

var log = logging.MustGetLogger("DBMergeMain")

func CheckError(err error) {
	if err != nil {
		log.Error(err)
	}
}
