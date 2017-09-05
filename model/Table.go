package model

import "github.com/oweissbarth/DBMerge/utils"

type Table struct {
	Name       string
	Database   Database
	PrimaryKey string
	Columns    []string
}

func (t *Table) GetPrimaryKey() string {
	if t.PrimaryKey != "" {
		return t.PrimaryKey
	}
	err := Con.QueryRow("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA='" + t.Database.Name + "' AND TABLE_NAME='" + t.Name + "' AND COLUMN_KEY='PRI'").Scan(&t.PrimaryKey)
	utils.CheckError(err)
	log.Debug("Primary key of " + t.Name + " is " + t.PrimaryKey)
	return t.PrimaryKey
}

func (t *Table) GetColumns() []string {
	if t.Columns != nil {
		log.Debug("Columns are cached.")
		return t.Columns
	}
	log.Debug("Columns are not cached. Querying.")

	rows, err := Con.Query("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA='" + t.Database.Name + "' AND TABLE_NAME='" + t.Name + "'")
	utils.CheckError(err)

	defer rows.Close()

	var column string

	for rows.Next() {
		err = rows.Scan(&column)
		utils.CheckError(err)
		t.Columns = append(t.Columns, column)
	}
	return t.Columns
}
