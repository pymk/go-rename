package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Parse arguments to get the directory path
	args := os.Args[1:]
	dir, err := parseArgs(args)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// List of file extensions to consider
	extensions := []string{"jpeg", "jpg", "png", "JPEG", "JPG", "PNG"}

	// List of files in the directory with specified extensions
	fNames, err := listFiles(dir, extensions)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// Inform user if no supported extensions found
	if len(fNames) == 0 {
		fmt.Printf("No files found in %s with the following extensions: %s\n", dir, extensions)
		return
	}

	// Determine the length of the longest file name for alignment
	padLen := longestChar(fNames)

	// Iterate over file names, calculate SHA hash, and print formatted output
	for _, f := range fNames {
		shaName := makeSha(f)
		ext := getExtension(f)
		cleanName := shaName + "." + ext
		fmt.Printf("%-*s %s\n", padLen, f, cleanName)
	}
}

func makeSha(name string) string {
	// SHA1 hash
	h := sha1.New()
	h.Write([]byte(name))
	sha := h.Sum(nil)

	// Hexadecimal conversion
	shaStr := hex.EncodeToString(sha)
	return shaStr
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

func parseArgs(args []string) (string, error) {
	var dirParsed string

	switch len(args) {
	case 0:
		dirParsed = "./"
	case 1:
		dirParsed = args[0]
	default:
		return "", errors.New("Usage: `program_name` or `program_name directory_path`")
	}

	return dirParsed, nil
}

func getExtension(name string) string {
	return strings.ToLower(name[strings.LastIndex(name, ".")+1:])
}

func listFiles(path string, extensions []string) ([]string, error) {
	ls, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("%v\n", err)
	}

	var fileNames []string
FileLoop:
	for _, f := range ls {
		name := f.Name()
		ext := getExtension(name)
		for _, e := range extensions {
			if ext == e {
				fileNames = append(fileNames, name)
				continue FileLoop
			}
		}
	}

	return fileNames, nil
}
