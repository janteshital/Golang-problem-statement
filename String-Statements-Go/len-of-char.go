package main

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Count the Chars
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func findLen() {
	// When you run it with "İNANÇ", it should return 5 not 7.

	name := "İNANÇ"
	length := utf8.RuneCountInString(name)
	fmt.Println(length)
}