package main

import (
	"fmt"

	"github.com/harry1453/go-common-file-dialog/cfd"
)

// executeFolderMode handles picking a folder and showing what you picked
func executeFolderMode() {
	// let's open a dialog to pick a folder
	dialog, err := cfd.NewSelectFolderDialog(cfd.DialogConfig{
		Title: "Select a Folder",
		Role:  "unethicalFolderPick",
	})
	handleError(err, false)

	// showing the dialog
	err = dialog.Show()
	handleError(err, false)

	// getting the result and showing it
	result, err := dialog.GetResult()
	handleError(err, true)

	fmt.Println(result)
}
