package model

import (
	"github.com/oweissbarth/DBMerge/utils"
)

/*Database is the representation of a database on a server*/
type Database struct {
	Name   string
	Tables []Table
}

/*GetTables returns a list of tables in the database*/
func (db Database) GetTables() []Table {
	if db.Tables != nil {
		return db.Tables
	}
	rows, err := Con.Query("SHOW TABLES IN " + db.Name)
	utils.CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var table string
		rows.Scan(&table)
		db.Tables = append(db.Tables, Table{Name: table, Database: db})
	}
	return db.Tables
}
