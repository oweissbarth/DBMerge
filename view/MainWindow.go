package view

import (
	"strconv"

	"github.com/gotk3/gotk3/gtk"
	"github.com/oweissbarth/DBMerge/control"
	"github.com/oweissbarth/DBMerge/model"
	"github.com/oweissbarth/DBMerge/utils"
)

func constructMainComparision(headDBName string, localDBName string, remoteDBName string, targetDBName string) {
	log.Info("Setting up main comparision")

	localDB := model.Database{Name: localDBName}
	headDB := model.Database{Name: headDBName}
	remoteDB := model.Database{Name: remoteDBName}
	//targetDB := model.Database{Name: targetDBName}

	localTables := localDB.GetTables()
	headTables := headDB.GetTables()
	remoteTables := remoteDB.GetTables()

	obj, err := builder.GetObject("mainContainer")
	utils.CheckError(err)

	widget := obj.(*gtk.ListBox)

	for i := range localTables {
		if localTables[i].Name != headTables[i].Name || localTables[i].Name != remoteTables[i].Name {
			log.Error("Table names don't match:", localTables[i].Name, headTables[i].Name, remoteTables[i].Name)
			return
		}

		localdiff := control.GetDiff(localTables[i], headTables[i])
		remotediff := control.GetDiff(remoteTables[i], headTables[i])
		patch := control.Merge(localdiff, remotediff)

		expander, err := gtk.ExpanderNew(localTables[0].Name)
		utils.CheckError(err)

		expander.SetLabel(localTables[i].Name)
		container, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 15)
		utils.CheckError(err)

		expander.Add(container)

		conflictLabel, err := gtk.LabelNew("Conflicts")
		utils.CheckError(err)
		container.Add(conflictLabel)

		conflictContainer, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 15)
		utils.CheckError(err)
		for j := range patch.Conflicts {
			conflict, err := gtk.LabelNew(strconv.Itoa(patch.Conflicts[j].PrimaryKey))
			utils.CheckError(err)
			conflictContainer.Add(conflict)
		}
		container.Add(conflictContainer)

		additionsLabel, err := gtk.LabelNew("Additions")
		utils.CheckError(err)
		container.Add(additionsLabel)

		additionContainer, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 15)
		utils.CheckError(err)
		for j := range patch.Additions {
			addition, err := gtk.LabelNew(strconv.Itoa(patch.Additions[j].PrimaryKey))
			utils.CheckError(err)
			additionContainer.Add(addition)
		}

		container.Add(additionContainer)

		deletionsLabel, err := gtk.LabelNew("Deletions")
		utils.CheckError(err)
		container.Add(deletionsLabel)

		deletionsContainer, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 15)
		utils.CheckError(err)
		for j := range patch.Deletions {
			deletion, err := gtk.LabelNew(strconv.Itoa(patch.Deletions[j].PrimaryKey))
			utils.CheckError(err)
			deletionsContainer.Add(deletion)
		}

		container.Add(deletionsContainer)

		modificationsLabel, err := gtk.LabelNew("Modifications")
		utils.CheckError(err)
		container.Add(modificationsLabel)

		modificationsContainer, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 15)
		utils.CheckError(err)
		for j := range patch.Modifications {
			modification, err := gtk.LabelNew(strconv.Itoa(patch.Modifications[j].PrimaryKey))
			utils.CheckError(err)
			modificationsContainer.Add(modification)
		}

		container.Add(modificationsContainer)

		widget.Add(expander)

	}
	widget.ShowAll()

}
