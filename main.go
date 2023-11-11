package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/harry1453/go-common-file-dialog/cfd"
)

// constants for the modes and exit codes
const (
	modeFile   = "file"
	modeFolder = "folder"
	exitErr    = 1
	exitCancel = 2
)

// flags for command-line arguments
var (
	mode = flag.String("mode", "", "Mode of operation: 'file' or 'folder'")
	ext  = flag.String("ext", "", "File extension filter (used in 'file' mode)")
)

// handleError deals with errors and exits if needed
func handleError(err error, exitOnErr bool) {
	if err != nil {
		if err == cfd.ErrorCancelled {
			fmt.Println("Operation cancelled by the user.")
			os.Exit(exitCancel)
		} else {
			fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
		}
		if exitOnErr {
			os.Exit(exitErr)
		}
	}
}

// showHelp displays help info about how to use this program
func showHelp() {
	fmt.Println("Usage of picker.exe:")
	flag.PrintDefaults()
	fmt.Println("\nExample:")
	fmt.Println("  picker.exe -mode file -ext .txt")
	fmt.Println("  picker.exe -mode folder")
}

// cleanExtensions takes your file extensions and makes them nice and tidy
func cleanExtensions(rawExt string) []string {
	var extensions []string
	for _, ext := range strings.Split(rawExt, ",") {
		ext = strings.TrimSpace(ext)
		if ext != "" {
			if !strings.HasPrefix(ext, ".") {
				ext = "." + ext
			}
			extensions = append(extensions, ext)
		}
	}
	return extensions
}

// validateMode checks if the mode is either 'file' or 'folder'
func validateMode(m string) bool {
	return m == modeFile || m == modeFolder
}

// validateExtensions makes sure your file extensions look good
func validateExtensions(exts []string) bool {
	for _, ext := range exts {
		if !strings.HasPrefix(ext, ".") || len(ext) < 2 {
			return false
		}
	}
	return true
}

// main is where everything starts
func main() {
	flag.Usage = showHelp
	flag.Parse()

	// need to pick a mode to get going
	if *mode == "" {
		showHelp()
		os.Exit(exitErr)
	}

	// check if the mode is valid
	if !validateMode(strings.ToLower(*mode)) {
		fmt.Println("Invalid mode provided.")
		showHelp()
		os.Exit(exitErr)
	}

	// clean up and validate extensions if provided
	cleanedExtensions := cleanExtensions(*ext)
	if len(cleanedExtensions) > 0 && !validateExtensions(cleanedExtensions) {
		fmt.Println("Invalid file extension format.")
		os.Exit(exitErr)
	}

	// decide what to do based on the mode
	switch strings.ToLower(*mode) {
	case modeFile:
		executeFileMode(*ext)
	case modeFolder:
		executeFolderMode()
	default:
		fmt.Println("Invalid mode provided.")
		showHelp()
		os.Exit(exitErr)
	}
}
