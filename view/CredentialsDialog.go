package view

import (
	"github.com/gotk3/gotk3/gtk"
	logging "github.com/op/go-logging"
	"github.com/oweissbarth/DBMerge/control"
	"github.com/oweissbarth/DBMerge/utils"
)

var log = logging.MustGetLogger("DBMergeMain")

func connectToServer() {
	usernameWidget, err := builder.GetObject("userNameField")
	utils.CheckError(err)

	passwordWidget, err := builder.GetObject("passwordField")
	utils.CheckError(err)

	hostnameWidget, err := builder.GetObject("hostnameField")
	utils.CheckError(err)

	portWidget, err := builder.GetObject("portField")
	utils.CheckError(err)

	usernameField := usernameWidget.(*gtk.Entry)
	passwordField := passwordWidget.(*gtk.Entry)
	hostnameField := hostnameWidget.(*gtk.Entry)
	portField := portWidget.(*gtk.Entry)

	username, err := usernameField.GetText()
	utils.CheckError(err)
	password, err := passwordField.GetText()
	utils.CheckError(err)
	hostname, err := hostnameField.GetText()
	utils.CheckError(err)
	port, err := portField.GetText()
	utils.CheckError(err)

	control.ConnectToServer(username, password, hostname, port)

	constructDatabaseSelection()

	destroyCredentials()
}

func destroyCredentials() {
	obj, err := builder.GetObject("credentialsDialog")
	utils.CheckError(err)

	dialog := obj.(*gtk.Dialog)
	dialog.Destroy()

}

func constructCredentials() {
	obj, err := builder.GetObject("credentialsDialog")
	utils.CheckError(err)

	dialog, ok := obj.(*gtk.Dialog)
	if !ok {
		log.Error("object returned from glade file could not be casted to Dialog")
	}

	dialog.ShowAll()
}
