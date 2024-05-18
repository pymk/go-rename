package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Parse the command-line flags.
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 || len(args) > 1 {
		pkgPath := os.Args[0]
		pkgName := filepath.Base(pkgPath)
		fmt.Fprintf(os.Stderr, "Usage: %s /path/to/file\n", pkgName)
		os.Exit(1)
	}

	dirPath := args[0]

	// List of file extensions to consider
	extensions := map[string]struct{}{
		"jpeg": {}, "jpg": {}, "png": {}, "JPEG": {}, "JPG": {}, "PNG": {},
	}

	fileNames, err := listFiles(dirPath, extensions)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	if len(fileNames) == 0 {
		fmt.Printf("No files found in %s with the following extensions: %v\n", dirPath, getKeys(extensions))
		return
	}

	// Determine the length of the longest file name for alignment
	padLen := longestChar(fileNames)

	fmt.Printf("%-*s %s\n", padLen, "Original", "New")
	fmt.Println(strings.Repeat("-", padLen-2) + strings.Repeat(" ", 3) + strings.Repeat("-", padLen-2))

	if err := renameFiles(dirPath, fileNames, padLen); err != nil {
		fmt.Printf("could not rename: %v\n", err)
	}
}

func renameFiles(dirPath string, fileNames []string, padLen int) error {
	for _, oldName := range fileNames {
		oldFullPath := filepath.Join(dirPath, oldName)
		shaName := makeSha(oldName)
		ext := getExtension(oldName)
		newName := shaName + "." + ext
		newFullPath := filepath.Join(dirPath, newName)

		if err := os.Rename(oldFullPath, newFullPath); err != nil {
			return fmt.Errorf("could not rename %s: %v", oldFullPath, err)
		} else {
			fmt.Printf("%-*s %s\n", padLen, oldName, newName)
		}
	}
	return nil
}

func longestChar(files []string) int {
	maxLen := 0
	for _, f := range files {
		fnameLen := len(f)
		if fnameLen > maxLen {
			maxLen = fnameLen
		}
	}

	return maxLen + 2
}
