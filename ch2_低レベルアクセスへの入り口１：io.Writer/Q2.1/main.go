package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("A2_1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%d %s %f", 1, "Hello", 1.0)
}
