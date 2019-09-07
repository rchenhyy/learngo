package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1) // defer：先进后出顺序
	defer fmt.Println(2)
	fmt.Println(3)
}

func main() {
	tryDefer()

	file, err := os.Create("abc.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer, "abc")
	fmt.Fprintln(writer, "efg")
	fmt.Fprintln(writer, "hyy")

}
