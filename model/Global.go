package model

import (
	"database/sql"

	logging "github.com/op/go-logging"
)

/*Con is the global connection to the mysql server*/
var Con *sql.DB

var log = logging.MustGetLogger("DBMergeMain")
