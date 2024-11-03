package main

import (
	"fmt"
	"io"
	"strings"
)

func reader(){
	r := strings.NewReader("Hello, Reader!")

	// makes a byte slice of len 8
	b := make([]byte, 8)

	// continuosly reads the file 8 bytes at a time until there's nothing left to read and the reader returns an EOF(end of file) error
	for {
		n, err := r.Read(b)

		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}