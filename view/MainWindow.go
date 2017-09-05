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
		diff := control.GetDiff(localTables[i], headTables[i])
		log.Info(strconv.Itoa(len(diff.Additions)) + " Additions in " + localTables[i].Name)
		log.Info(strconv.Itoa(len(diff.Deletions)) + " Deletions in " + localTables[i].Name)
		log.Info(strconv.Itoa(len(diff.Modifications)) + " Modifications in " + localTables[i].Name)

	}
}
