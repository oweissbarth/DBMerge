package model

import (
	"database/sql"

	logging "github.com/op/go-logging"
)

var Con *sql.DB

var log = logging.MustGetLogger("DBMergeMain")
