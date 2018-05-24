package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		if 'A' <= b[i] && b[i] <= 'Z' {
			b[i] = (b[i] + 13)
			if b[i] > 'Z' {
				b[i] -= 26
			}
		} else if 'a' <= b[i] && b[i] <= 'z' {
			b[i] = (b[i] + 13)
			if b[i] > 'z' {
				b[i] -= 26
			}
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
