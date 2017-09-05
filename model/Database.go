package model

import (
	"github.com/oweissbarth/DBMerge/utils"
)

/*Database is the representation of a database on a server*/
type Database struct {
	Name   string
	Tables []Table
}

func (db Database) GetTables() []Table {
	rows, err := Con.Query("SHOW TABLES IN " + db.Name)
	utils.CheckError(err)

	for rows.Next() {
		var table string
		rows.Scan(&table)
		db.Tables = append(db.Tables, Table{Name: table, Database: db})
	}
	return db.Tables
}
