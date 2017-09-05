package model

import "github.com/oweissbarth/DBMerge/utils"

/*Table represents a table*/
type Table struct {
	Name       string
	Database   Database
	PrimaryKey string
	Columns    []string
}

/*GetPrimaryKey returns the primary key of the table*/
func (t *Table) GetPrimaryKey() string {
	if t.PrimaryKey != "" {
		return t.PrimaryKey
	}
	err := Con.QueryRow("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA='" + t.Database.Name + "' AND TABLE_NAME='" + t.Name + "' AND COLUMN_KEY='PRI'").Scan(&t.PrimaryKey)
	utils.CheckError(err)
	return t.PrimaryKey
}

/*GetColumns returns a list containing the column names as strings*/
func (t *Table) GetColumns() []string {
	if t.Columns != nil {
		return t.Columns
	}

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
