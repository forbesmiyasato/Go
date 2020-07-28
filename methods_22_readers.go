package main

import "golang.org/x/tour/reader"

type MyReader struct{}

//Implement the reader interface
func (reader MyReader) Read(x []byte) (int, error) {
	for i := 0; i < 10; i++ {
		x[i] = 'A'
	}
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
