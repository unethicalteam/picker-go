package main

import (
	"fmt"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
)

// executeFileMode is all about picking a file and showing your choice
func executeFileMode(ext string) {
	// preparing filters based on the extensions you want
	fileFilters := []cfd.FileFilter{}
	for _, ext := range cleanExtensions(ext) {
		fileFilters = append(fileFilters, cfd.FileFilter{DisplayName: ext[1:] + " Files", Pattern: "*" + ext})
	}

	// here we open the file dialog
	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title:       "Select a File",
		Role:        "unethicalFilePick",
		FileFilters: fileFilters,
	})
	handleError(err, true)

	// and here we show what you picked
	fmt.Println(result)
}
