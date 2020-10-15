package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"

func (reader spaceEraser) Read(p []byte) (int, error) { //parcours list for tha amount of chara read
	n, err := reader.r.Read(p)
	charact := 0 //counter for character != space
	for i := 0; i<n; i++{
		if p[i] != 32{ //test is charact is a space (32 = ASCII space)
			p[charact] = p[i]
			charact++
		}
	}
	return charact, err
}

