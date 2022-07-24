
package main

import "fmt"

func main() {
	fmt.Println("parenthesis validation")
	str := "{[[]]}"
	res := checkValidParenthesis(str)
	fmt.Println(res)

}

func checkValidParenthesis(str string) bool {
	arr := []string{}
	valid := true
	for _, r := range str {
		s := string(r)
		if s == "[" || s == "{" || s == "(" {
			arr = append(arr, s)
		} else if s == "]" {
			e := arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			if e != "[" {
				valid = false
				break
			}
		} else if s == "}" {
			e := arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			if e != "{" {
				valid = false
				break
			}
		} else if s == ")" {
			e := arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			if e != "(" {
				valid = false
				break
			}
		}
	}
	return valid

}
