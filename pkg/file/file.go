package file

import (
	"fmt"
	"os"
	"strings"
)

func getKeys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getExtension(name string) string {
	return strings.ToLower(name[strings.LastIndex(name, ".")+1:])
}

func listFiles(path string, extensions map[string]struct{}) ([]string, error) {
	ls, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("%v\n", err)
	}

	var fileNames []string
	for _, f := range ls {
		name := f.Name()
		ext := getExtension(name)
		if _, exists := extensions[ext]; exists {
			fileNames = append(fileNames, name)
		}
	}

	return fileNames, nil
}
