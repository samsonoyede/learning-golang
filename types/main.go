package main

import (
	"fmt"
	"os"
)

// strings are immutable

func main() {
	var sum float64
	var n int
	for {
		var value float64

		// Reading from command line
		_, err := fmt.Fscanln(os.Stdin, &value)
		// if there is an error break the programme
		if err != nil {
			break
		}
		sum += value
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "No Values")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/float64(n))
}
