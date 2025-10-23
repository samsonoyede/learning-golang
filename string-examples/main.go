package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// s := "Say Hello"
	// // Get length of a string
	// // lengthOfS := len(s)

	// // substring
	// getValueOfSay := s[:3]
	// fmt.Println(getValueOfSay)
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	// initializing two variables at the same time
	// This should only be done when both variables are the output of a function or expression

	old, new := os.Args[1], os.Args[2]
	// Scans one line from the input
	scan := bufio.NewScanner(os.Stdin)

	// Scan() scans the input and if it succeeds
	for scan.Scan() {
		// splits input and does not include the delimiter
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
}

// runes are bytes in golang
// the length of a string is the number bytes required to represent the number of unicode characters
