package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
)

const (
	modeFile   = "file"
	modeFolder = "folder"
	credit     = "developed by unethical.\ncopyright 2023\n"
)

func showError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("Usage: picker.exe <mode> [ext for file mode]")
	fmt.Println("Modes:")
	fmt.Println("\tfile [ext]  - Opens a dialog to select a file with the optional extension.")
	fmt.Println("\tfolder      - Opens a dialog to select a folder.")
}

func main() {
	fmt.Println(credit)
	if len(os.Args) < 2 {
		showHelp()
		os.Exit(1)
	}

	mode := strings.ToLower(os.Args[1])

	switch mode {
	case modeFile:
		fileFilters := []cfd.FileFilter{}
		if len(os.Args) == 3 {
			ext := os.Args[2]
			if !strings.HasPrefix(ext, ".") {
				ext = "." + ext
			}
			fileFilters = append(fileFilters, cfd.FileFilter{DisplayName: ext[1:] + " Files", Pattern: "*" + ext})
		}

		result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
			Title:       "Select a File",
			Role:        "unethicalFilePick",
			FileFilters: fileFilters,
		})
		showError(err)
		fmt.Println(result)

	case modeFolder:
		dialog, err := cfd.NewSelectFolderDialog(cfd.DialogConfig{
			Title: "Select a Folder",
			Role:  "unethicalFolderPick",
		})
		showError(err)
		err = dialog.Show()
		showError(err)
		result, err := dialog.GetResult()
		showError(err)
		fmt.Println(result)

	default:
		fmt.Println("Invalid mode provided.")
		showHelp()
		os.Exit(1)
	}
}
