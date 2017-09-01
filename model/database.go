package model

import (
	"database/sql"
)

type Database struct {
	sql.DB
	name   string
	tables []string
}

func (db Database) getTables() {
	rows, err := db.Query("SHOW TABLES")
	checkError(err)

	for rows.Next() {
		var table string
		rows.Scan(&table)
		db.tables = append(db.tables, table)
	}
}

func open(user string, password string, host string, db string) *sql.DB {
	database, err := sql.Open("mysql", user+":"+password+"@"+host+"/"+db)
	checkError(err)
	return database
}
