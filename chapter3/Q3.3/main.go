package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

type MyZipWriter struct {
	zw *zip.Writer
}

func NewMyZipWriter(zw *zip.Writer) *MyZipWriter {
	return &MyZipWriter{zw: zw}
}

func (mzw *MyZipWriter) Write(name string, r *strings.Reader) (io.Writer, error) {
	file, err := mzw.zw.Create(name)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(file, r)
	if err != nil {
		return nil, err
	}
	return file, err
}

func main() {
	file, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	myZipWriter := NewMyZipWriter(zipWriter)
	_, err = myZipWriter.Write("newfile.txt", strings.NewReader("this is a test"))
	if err != nil {
		panic(err)
	}
}
