package main

import (
	"archive/zip"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	file, err := zipWriter.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	io.WriteString(file, "Q3.4 sample")
	zipWriter.Flush()
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachement; filename=ascii_sample.zip")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
