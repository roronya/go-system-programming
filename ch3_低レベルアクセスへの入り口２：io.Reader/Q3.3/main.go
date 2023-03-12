package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	zipFile, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	file, err := zipWriter.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(file, strings.NewReader("this is a sample text"))
	zipWriter.Flush()
}
