package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	gzw := gzip.NewWriter(w)
	source := map[string]string{
		"Hello": "World",
	}
	// write code here
	mw := io.MultiWriter(gzw, os.Stdout)
	encoder := json.NewEncoder(mw)
	encoder.SetIndent("", "	")
	err := encoder.Encode(source)
	if err != nil {
		panic(err)
	}
	err = gzw.Flush()
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
