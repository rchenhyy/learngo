package container

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello 哈哈哈！"
	for i, b := range []byte(s) {
		fmt.Printf("(%d %X) ", i, b)
	}
	fmt.Println()
	for i, ch := range s { // ch: int32; ch is a rune
		fmt.Printf("(%d %X) ", i, ch) // i is the index of bytes
	}
	fmt.Println()

	fmt.Println("Byte count: ", len(s))                    // byte count, not decoded
	fmt.Println("Rune count: ", utf8.RuneCountInString(s)) // rune count, with utf8 decode

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	fmt.Println()
	for i, ch := range []rune(s) { // a new rune array
		fmt.Printf("(%d %c)", i, ch)
	}
}
