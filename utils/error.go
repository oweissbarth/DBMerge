package utils

import (
	"runtime"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("DBMergeMain")

func CheckError(err error) {
	_, file, line, _ := runtime.Caller(1)
	if err != nil {
		log.Error(file, line, err)
	}
}

func CompareSlices(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
