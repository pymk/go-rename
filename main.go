package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pymk/go-rename/cmd"
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

	cmd.Execute(dirPath)

}
