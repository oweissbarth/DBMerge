package view

import (
	"strconv"

	"github.com/oweissbarth/DBMerge/control"
	"github.com/oweissbarth/DBMerge/model"
)

func constructMainComparision(headDBName string, localDBName string, remoteDBName string, targetDBName string) {
	log.Info("Setting up main comparision")

	localDB := model.Database{Name: localDBName}
	headDB := model.Database{Name: headDBName}
	//remoteDB := model.Database{Name: remoteDBName}

	localTables := localDB.GetTables()
	headTables := headDB.GetTables()
	//remoteTables := remoteDB.GetTables()

	for i := range localTables {
		additions := control.GetAdditions(localTables[i], headTables[i])
		deletions := control.GetAdditions(headTables[i], localTables[i])

		log.Info(strconv.Itoa(len(additions)) + " Additions in " + localTables[i].Name)
		log.Info(strconv.Itoa(len(deletions)) + " Deletions in " + localTables[i].Name)

	}
}
