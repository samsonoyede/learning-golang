package main

import (
	"fmt"
	"hello"
	"os"
)

// 1: there is saying start from the second item in the string till the end
func main() {
	fmt.Println(hello.SpeakNames(os.Args[1:]))
}
