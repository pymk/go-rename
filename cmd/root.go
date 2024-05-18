package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pymk/go-rename/pkg/file"
	"github.com/pymk/go-rename/pkg/hasher"
)

func Execute(dirPath string) {
	// List of file extensions to consider
	extensions := map[string]struct{}{
		"jpeg": {}, "jpg": {}, "png": {}, "JPEG": {}, "JPG": {}, "PNG": {},
	}

	fileNames, err := file.ListFiles(dirPath, extensions)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	if len(fileNames) == 0 {
		fmt.Printf("No files found in %s with the following extensions: %v\n", dirPath, file.GetKeys(extensions))
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
		shaName := hasher.MakeSha(oldName)
		ext := file.GetExtension(oldName)
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
