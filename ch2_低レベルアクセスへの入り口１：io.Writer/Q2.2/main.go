package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("A2_2.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"column0", "column1", "column2"})
	writer.Write([]string{"record0", "record1", "record2"})
	writer.Flush()
	writer = csv.NewWriter(os.Stdout)
	writer.Write([]string{"column0", "column1", "column2"})
	writer.Write([]string{"record0", "record1", "record2"})
	writer.Flush()
}
