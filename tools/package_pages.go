// Prepares the generated site for GitHub Pages.
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func copyFile(sourcePath, destinationPath string, mode os.FileMode) {
	source, err := os.Open(sourcePath)
	check(err)
	defer source.Close()

	destination, err := os.OpenFile(destinationPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	check(err)
	defer destination.Close()

	_, err = io.Copy(destination, source)
	check(err)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run tools/package_pages.go DESTINATION")
		os.Exit(2)
	}

	destinationDir := os.Args[1]
	check(os.MkdirAll(destinationDir, 0755))
	existingEntries, err := os.ReadDir(destinationDir)
	check(err)
	if len(existingEntries) != 0 {
		panic(fmt.Sprintf("%s: destination must be empty", destinationDir))
	}

	entries, err := os.ReadDir("public")
	check(err)
	copied := 0
	for _, entry := range entries {
		if entry.IsDir() {
			panic(fmt.Sprintf("public/%s: directories are not supported", entry.Name()))
		}

		outputName := entry.Name()
		if filepath.Ext(outputName) == "" {
			outputName += ".html"
		}

		info, err := entry.Info()
		check(err)
		copyFile(filepath.Join("public", entry.Name()), filepath.Join(destinationDir, outputName), info.Mode())
		copied++
	}
	fmt.Printf("Prepared %d files for GitHub Pages.\n", copied)
}
