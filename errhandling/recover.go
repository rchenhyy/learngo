package main

import (
	"fmt"
)

func main() {
	defer func() {
		r := recover()
		if err, ok := r.(error); r != nil && ok {
			fmt.Printf("This is an error: %s\n", err.Error())
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("A new error"))
}
