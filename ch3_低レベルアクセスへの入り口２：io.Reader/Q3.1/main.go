package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		flag.Usage()
		os.Exit(1)
	}
	oldFileName := flag.Arg(0)
	newFileName := flag.Arg(1)
	oldFile, err := os.Open(oldFileName)
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()
	newFile, err := os.Create(newFileName)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	io.Copy(newFile, oldFile)
}
