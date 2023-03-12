package myio

import (
	"bytes"
	"io"
	"testing"
)

func TestCopyN(t *testing.T) {
	src := bytes.NewBufferString("this is a sample text.")
	dst := bytes.NewBuffer([]byte{})
	CopyN(dst, src, 10)

	got, _ := io.ReadAll(dst)
	if string(got) != "this is a " {
		t.Fatalf("want \"this is a \", but got %s\n", got)
	}

}
