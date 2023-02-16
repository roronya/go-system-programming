package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	// 最初の8バイトを飛ばす
	file.Seek(8, 0)
	var offset int64 = 8
	for {
		p, _ := file.Seek(0, 1)
		fmt.Println(p)
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		p, _ = file.Seek(0, 1)
		fmt.Println(p)
		chunks = append(chunks, io.NewSectionReader(file, offset, 4+int64(length)+4+4))
		p, _ = file.Seek(0, 1)
		fmt.Println(p)

		offset, _ = file.Seek(4+int64(length)+4, 1)
	}
	return chunks
}

func main() {
	file, err := os.Open("PNG_transparency_demonstration_1.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}