package main

import (
	"io"
	"os"
)

type ReversStringReader struct {
	input string
}

func (rsr *ReversStringReader) Read(p []byte) (int, error) {
	n := copy(p, []byte(rsr.input))
	return n, io.EOF
}
func NewReversStringReader(input string) *ReversStringReader {
	var revS string
	for i := len([]byte(input)) - 1; i >= 0; i-- {
		revS += string([]byte(input)[i])
	}
	return &ReversStringReader{input: revS}
}
func main() {
	ad := NewReversStringReader("Reverse this String")
	io.Copy(os.Stdout, ad)
}
