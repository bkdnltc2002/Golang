package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	fmt.Printf("\n%v %v\n", n, err)
	const A byte = 'A'
	const a byte = 'a'
	for i, x := range b {
		if x == 0 {
			n = i
			break
		}
		switch {
		case A <= x && x < a:
			tmp := x - A
			b[i] = A + ((tmp + 13) % 26)
		case a <= x && x < a+26:
			tmp := x - a
			b[i] = a + ((tmp + 13) % 26)
		}
	}
	fmt.Println("About to return n: ", n, "buffer is:", string(b))
	return n, err
}

func main() {
	// r := strings.NewReader("Hello, Reader! ")

	// b := make([]byte, 8)
	// for {
	// 	n,err := r.Read(b)
	// 	fmt.Printf("n = %v err = %v\n",n,err,b)
	// 	fmt.Printf("b[:n] = %q\n",b[:n])
	// 	if err == io.EOF{
	// 		break
	// 	}
	// }
	s := strings.NewReader("LBH penxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
