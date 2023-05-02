package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := stringutil.Reverse(input)
	doubleRev := stringutil.Reverse(rev)

	fmt.Printf("Original: %q\n", input)
	fmt.Printf("Reversed: %q\n", rev)
	fmt.Printf("Reversed again: %q\n", doubleRev)
}
