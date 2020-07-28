package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13 (x byte) byte {
	if (x >= 65 && x <= 90) {
		x+=13
		if x > 90 { x-=26 }
	}
	if (x >= 97 && x <= 121) {
		x+=13
		if x > 121 {x-=26}
	}
	return x
}

func (reader rot13Reader) Read(x []byte) (int, error) {
	len, err := reader.r.Read(x)

	for index := range x {
		x[index] = rot13(x[index])
	}

	return len, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
