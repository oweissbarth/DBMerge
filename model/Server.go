package model

import (
	"database/sql"
	"strconv"

	"github.com/oweissbarth/DBMerge/utils"
)

/*Server represents a mysql server. It holds login information as well as a list of the databases*/
type Server struct {
	Hostname  string
	Port      int
	Username  string
	Password  string
	Databases []Database
}

/*Connect authentificates with the mysql server*/
func (s Server) Connect() {
	if Con != nil {
		log.Error("already connected to a server")
	}

	var err error
	Con, err = sql.Open("mysql", s.Username+":"+s.Password+"@tcp("+s.Hostname+":"+strconv.Itoa(s.Port)+")/")
	utils.CheckError(err)

	log.Info("Connected to the server")
}

/*GetDatabases retrieves all available databases from the mysql server*/
func (s Server) GetDatabases() []Database {
	if Con == nil {
		log.Error("Could not retrieve database list as we are not connected to a server")
		return nil
	}

	if s.Databases != nil {
		return s.Databases
	}

	rows, err := Con.Query("SHOW DATABASES")
	utils.CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var dbName string
		rows.Scan(&dbName)
		s.Databases = append(s.Databases, Database{Name: dbName})
	}
	return s.Databases
}
