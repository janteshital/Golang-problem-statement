package main


import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------

func main() {

	//-1 is added as args[0] contains the file name
	count := len(os.Args)-1

	fmt.Printf("There are %d names.\n", count)
}