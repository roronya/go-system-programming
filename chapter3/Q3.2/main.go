package main

import (
	"crypto/rand"
	"os"
)

func main() {
	p := make([]byte, 1024)
	_, err := rand.Reader.Read(p)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(p)
}
