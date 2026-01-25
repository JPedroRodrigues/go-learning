package main

import (
	"io"
	"os"
	"strings"
)


type rot13Reader struct {
	r io.Reader
}

func (rotR rot13Reader) Read(b []byte) (int, error) {
	readr, err := rotR.r.Read(b)

	if err != nil {
		return readr, err
	}

	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphaSize := len(alphabetUpper)

	for i := range(readr) {
		if idx := strings.Index(alphabetUpper, string(b[i])); idx != -1 {
			b[i] = alphabetUpper[(idx + 13) % alphaSize]
		} else if idx := strings.Index(alphabetLower, string(b[i])); idx != -1 {
			b[i] = alphabetLower[(idx + 13) % alphaSize]
		}
	}

	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}