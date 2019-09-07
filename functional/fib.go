package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(bytes []byte) (int, error) {
	next := g()
	if next > 1000000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(bytes) // incorrect if bytes is too small
}

func main3() {
	f := fibonacci()
	f()
	f()
	f()
	f()
	f()

	scanner := bufio.NewScanner(fibonacci())
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
