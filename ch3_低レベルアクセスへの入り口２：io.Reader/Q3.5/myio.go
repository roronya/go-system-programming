package myio

import (
	"io"
)

func CopyN(dest io.Writer, src io.Reader, length int) {
	reader := io.LimitReader(src, int64(length))
	io.Copy(dest, reader)
}
