package view

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/oweissbarth/DBMerge/utils"
)

var builder *gtk.Builder

/*Construct reads in the ui description file and sets up the main window*/
func Construct() {
	gtk.Init(nil)

	var err error
	builder, err = gtk.BuilderNew()
	utils.CheckError(err)
	err = builder.AddFromFile("view/MainWindow.glade")
	utils.CheckError(err)

	obj, err := builder.GetObject("window")
	utils.CheckError(err)

	win, ok := obj.(*gtk.Window)

	if !ok {
		log.Error("object returned from glade file could not be casted to Window")
	}

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.ShowAll()

	connectSignals(builder)

	constructCredentials()

	gtk.Main()
}

func connectSignals(builder *gtk.Builder) {
	builder.ConnectSignals(map[string]interface{}{
		"connectButtonHandler": connectToServer,
		"compareButtonHandler": handleCompareButton,
	})
}
