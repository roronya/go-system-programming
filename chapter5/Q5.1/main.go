package main

import (
	"io"
	"os"
)

func main() {
	in, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	out, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(out, in)
	if err != nil {
		panic(err)
	}
}
