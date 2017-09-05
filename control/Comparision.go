package control

import (
	m "github.com/oweissbarth/DBMerge/model"
	"github.com/oweissbarth/DBMerge/utils"
)

/*GetAdditions returns a list of a rows in A that are not in B*/
func GetAdditions(tableA m.Table, tableB m.Table) []m.Addition {
	columnsA := tableA.GetColumns()
	columnsB := tableB.GetColumns()

	if !utils.CompareSlices(columnsA, columnsB) {
		log.Error("The Schema of " + tableA.Database.Name + "." + tableA.Name + " and " + tableB.Database.Name + "." + tableB.Name + " are not the same ")
		return nil
	}

	var same string

	for i, col := range columnsA {
		if i == 0 {
			same = same + "A." + col + "=" + col
		} else {
			same = same + " AND A." + col + "=" + col
		}
	}

	rows, err := m.Con.Query("SELECT " + tableA.GetPrimaryKey() + " FROM " + tableA.Database.Name + "." + tableA.Name + " AS A WHERE NOT EXISTS(SELECT * FROM " + tableB.Database.Name + "." + tableB.Name + " WHERE " + same + ")")
	utils.CheckError(err)

	var additions []m.Addition
	var addition int

	for rows.Next() {
		rows.Scan(&addition)
		additions = append(additions, m.Addition{addition})
	}
	return additions
}
