package control

import (
	"strings"

	m "github.com/oweissbarth/DBMerge/model"
	"github.com/oweissbarth/DBMerge/utils"
)

/*GetAdditions returns a list of a rows in A that are not in B*/
func GetAdditions(tableA m.Table, tableB m.Table) []m.Addition {
	if !tableA.IsCompatibleWith(&tableB) {
		return nil
	}

	columnsA := tableA.GetColumns()

	var same string

	for i, col := range columnsA {
		if i == 0 {
			same = same + "A." + col + "=" + col
		} else {
			same = same + " AND A." + col + "=" + col
		}
	}

	rows, err := m.Con.Query("SELECT " + tableA.GetPrimaryKey() + ", CONCAT_WS(', ', " + strings.Join(columnsA, ",") + ")" + " FROM " + tableA.Database.Name + "." + tableA.Name + " AS A WHERE NOT EXISTS(SELECT * FROM " + tableB.Database.Name + "." + tableB.Name + " WHERE " + same + ")")
	utils.CheckError(err)

	var additions []m.Addition
	var addition int
	var content string

	for rows.Next() {
		rows.Scan(&addition, &content)
		additions = append(additions, m.Addition{PrimaryKey: addition, Content: content, Origin: &tableA})
	}
	return additions
}

/*GetDiff return a full differential between two tables*/
func GetDiff(tableA m.Table, tableB m.Table) m.Differential {
	additions := GetAdditions(tableA, tableB)
	inverseAdditions := GetAdditions(tableB, tableA)

	var deletions []m.Deletion
	for _, d := range inverseAdditions {
		deletions = append(deletions, m.Deletion{d.PrimaryKey, d.Content, d.Origin})
	}

	// Find modifications and remove those from additions and deletions
	var modifications []m.Modification

	for i := 0; i < len(additions); i++ {
		a := additions[i]
		for j, d := range deletions {
			if a.PrimaryKey == d.PrimaryKey {
				modifications = append(modifications, m.Modification{PrimaryKey: a.PrimaryKey})
				additions = append(additions[:i], additions[i+1:]...)
				deletions = append(deletions[:j], deletions[j+1:]...)
				i--
				break
			}
		}
	}
	return m.Differential{
		Additions:     additions,
		Deletions:     deletions,
		Modifications: modifications}
}
