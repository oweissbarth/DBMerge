package model

import (
	"strconv"

	"github.com/oweissbarth/DBMerge/utils"
)

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

func (t *Table) String() string {
	return "<Table: " + t.Name + ">"
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

/*IsCompatibleWith returns true the given table has the same schema as this table*/
func (t *Table) IsCompatibleWith(t2 *Table) bool {
	columns1 := t.GetColumns()
	columns2 := t2.GetColumns()

	diff := utils.CompareSlices(columns1, columns2)
	if diff != -1 {
		log.Error("The Schema of " + t.Database.Name + "." + t.Name + " and " + t2.Database.Name + "." + t2.Name + " are not the same.")
		var cA string
		if len(columns1) < diff+1 {
			cA = "out of range"
		} else {
			cA = columns1[diff]
		}

		var cB string
		if len(columns2) < diff+1 {
			cB = "out of range"
		} else {
			cB = columns2[diff]
		}
		log.Error("Column " + strconv.Itoa(diff+1) + " in " + t.Database.Name + " is : " + cA + " while in " + t2.Database.Name + " it is : " + cB)
		return false
	}
	return true
}
