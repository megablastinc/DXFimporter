package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/* File management */
func getFiles(path string, ext string) (files []string) {
	root := path

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		sanitizedPath := strings.ToUpper(filepath.Ext(path))
		sanitizedExt := strings.ToUpper(ext)
		if sanitizedPath == sanitizedExt {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return files
}

func writeFile(content []string, file string) {
	fileHandle, _ := os.Create(file)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for _, line := range content {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}
