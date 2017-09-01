package view

import("github.com/gotk3/gotk3/gtk"
			"github.com/oweissbarth/DBMerge/control"
			"github.com/oweissbarth/DBMerge/utils")


func constructDatabaseSelection(){
	dbs := control.GetDatabases()

	dialogWidget, err := builder.GetObject("DatabaseSelectionDialog")
	utils.CheckError(err)

	dialog := dialogWidget.(*gtk.Dialog)


	headDBWidget, err := builder.GetObject("headDBField")
	utils.CheckError(err)
	remoteDBWidget, err := builder.GetObject("remoteDBField")
	utils.CheckError(err)
	localDBWidget, err := builder.GetObject("localDBField")
	utils.CheckError(err)
	targetDBWidget, err := builder.GetObject("targetDBField")
	utils.CheckError(err)


	headDBCombo := headDBWidget.(*gtk.ComboBoxText)
	remoteDBCombo := remoteDBWidget.(*gtk.ComboBoxText)
	localDBCombo := localDBWidget.(*gtk.ComboBoxText)
	targetDBCombo := targetDBWidget.(*gtk.ComboBoxText)

	for _, db := range dbs{
		headDBCombo.AppendText(db)
		remoteDBCombo.AppendText(db)
		localDBCombo.AppendText(db)
		targetDBCombo.AppendText(db)
	}

	dialog.ShowAll()
}
