package hello

import (
	"fmt"
	"strings"
)

// type is after the variable
// in .net its string name
func Speak(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// [] string is not an array but a slice. See difference at
// https://dev.to/dawkaka/go-arrays-and-slices-a-deep-dive-dp8
func SpeakNames(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}

	// Simple concatenation
	return "Hello, " + strings.Join(names, ", ") + "!"
}
