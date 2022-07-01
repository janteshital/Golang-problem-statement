package main

import "fmt"

func main() {

	s := "abaa"
	res := checkIfValid(s)
	fmt.Printf("is string valid: %t \n", res)
}

func checkIfValid(s string) bool {
	newmap := make(map[string]int)
	for _, r := range s {
		val, ok := newmap[string(r)]
		if ok {
			newmap[string(r)] = val + 1
		} else {

			newmap[string(r)] = 1
		}
	}
	fmt.Println(newmap)

	maxval := 0
	valid := true
	for _, val := range newmap {
		if val > maxval {
			maxval = val
		}
	}
	requiredchar := 0
	for _, val := range newmap {
		if val != maxval {
			requiredchar += (maxval - val)
			valid = false
		}
	}
	fmt.Printf("total required char in order to make string valid:  %d \n", requiredchar)
	return valid

}
