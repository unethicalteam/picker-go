package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/harry1453/go-common-file-dialog/cfd"
    "github.com/harry1453/go-common-file-dialog/cfdutil"
)

const (
    modeFile     = "file"
    modeFolder   = "folder"
    credit       = "developed by unethical.\ncopyright 2023\n"
    exitErr      = 1
    exitCancel   = 2
)

func handleError(err error, exitOnErr bool) {
    if err != nil {
        if err == cfd.ErrorCancelled {
            fmt.Println("Operation cancelled by the user.")
            os.Exit(exitCancel)
        }

        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        if exitOnErr {
            os.Exit(exitErr)
        }
    }
}

func showHelp() {
    fmt.Println("Usage: picker.exe <mode> [ext for file mode]")
    fmt.Println("Modes:")
    fmt.Println("\tfile [ext]  - Opens a dialog to select a file with the optional extension.")
    fmt.Println("\tfolder      - Opens a dialog to select a folder.")
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println(credit)
        showHelp()
        os.Exit(exitErr)
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
        handleError(err, true)
        
        fmt.Println(result)

    case modeFolder:
        dialog, err := cfd.NewSelectFolderDialog(cfd.DialogConfig{
            Title: "Select a Folder",
            Role:  "unethicalFolderPick",
        })
        handleError(err, false)

        err = dialog.Show()
        handleError(err, false)

        result, err := dialog.GetResult()
        handleError(err, true)
        
        fmt.Println(result)

    default:
        fmt.Println("Invalid mode provided.")
        showHelp()
        os.Exit(exitErr)
    }
}