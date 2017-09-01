package control

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	logging "github.com/op/go-logging"
	"github.com/oweissbarth/DBMerge/utils"
)

var log = logging.MustGetLogger("DBMergeMain")

var con *sql.DB = nil

func ConnectToServer(username string, password string, hostname string, port string) {
	if con != nil {
		log.Error("already connected to a server")
	}

	var err error
	con, err = sql.Open("mysql", username+":"+password+"@tcp("+hostname+":"+port+")/")
	utils.CheckError(err)

	log.Info("Connected to the server")
}

func GetDatabases() []string {
	if con == nil {
		log.Error("Could not retrieve database list as we are not connected to a server")
		return nil
	}

	rows, err := con.Query("SHOW DATABASES")
	utils.CheckError(err)

	defer rows.Close()

	dbs := []string{}

	for rows.Next() {
		var dbName string
		rows.Scan(&dbName)
		dbs = append(dbs, dbName)
	}
	return dbs
}
