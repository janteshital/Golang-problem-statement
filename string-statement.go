package main

import (
	"fmt"
)

func main() {
	x := "  Shital Jante  "
	x = checkIfHasSpace(x)
	fmt.Println("removed space", x)
}


func checkIfHasSpace(s string) string{
	for i, eachvar := range s{
		if eachvar == 32 {
			return checkIfHasSpace(s[:i]+s[i+1:])	
		}
	}
	return s
}
