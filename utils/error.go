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

func CompareSlices(a, b []string) int {
	if a == nil && b == nil {
		return -1
	}

	if a == nil || b == nil {
		return 0
	}

	var i int
	for i = 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			return i
		}
	}
	if len(a) != len(b) {
		return i
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
